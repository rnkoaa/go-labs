package store

import (
	"crypto/tls"
	"errors"
	"time"
)

// Backend represents a KV Store Backend
type Backend string

const (
	// CONSUL backend
	CONSUL Backend = "consul"
	// ETCD backend
	ETCD Backend = "etcd"
	// ZK backend
	ZK Backend = "zk"
	// BOLTDB backend
	BOLTDB Backend = "boltdb"
)

var (
	// ErrBackendNotSupported is thrown when the backend k/v store is not supported by libkv
	ErrBackendNotSupported = errors.New("backend storage not supported yet, please choose one of")
	// ErrCallNotSupported is thrown when a method is not implemented/supported by the current backend
	ErrCallNotSupported = errors.New("the current call is not supported with this backend")
	// ErrNotReachable is thrown when the API cannot be reached for issuing common store operations
	ErrNotReachable = errors.New("api not reachable")
	// ErrCannotLock is thrown when there is an error acquiring a lock on a key
	ErrCannotLock = errors.New("error acquiring the lock")
	// ErrKeyModified is thrown during an atomic operation if the index does not match the one in the store
	ErrKeyModified = errors.New("unable to complete atomic operation, key modified")
	// ErrKeyNotFound is thrown when the key is not found in the store during a Get operation
	ErrKeyNotFound = errors.New("key not found in store")
	// ErrPreviousNotSpecified is thrown when the previous value is not specified for an atomic operation
	ErrPreviousNotSpecified = errors.New("previous K/V pair should be provided for the Atomic operation")
	// ErrKeyExists is thrown when the previous value exists in the case of an AtomicPut
	ErrKeyExists = errors.New("previous K/V pair exists, cannot complete Atomic operation")
)

// Config contains the options for a storage client
type Config struct {
	ClientTLS         *ClientTLSConfig
	TLS               *tls.Config
	ConnectionTimeout time.Duration
	Bucket            string
	PersistConnection bool
	Username          string
	Password          string
}

// ClientTLSConfig contains data for a Client TLS configuration in the form
// the etcd client wants it.  Eventually we'll adapt it for ZK and Consul.
type ClientTLSConfig struct {
	CertFile   string
	KeyFile    string
	CACertFile string
}

// Store represents the backend K/V storage
// Each store should support every call listed
// here. Or it couldn't be implemented as a K/V
// backend for libkv
type Store interface {
	// Put a value at the specified key
	Put(key string, value []byte, options *WriteOptions) error

	// Get a value given its key
	Get(key string) (*KVPair, error)

	// Delete the value at the specified key
	Delete(key string) error

	// Verify if a Key exists in the store
	Exists(key string) (bool, error)

	// List the content of a given prefix
	List(directory string) ([]*KVPair, error)

	// Atomic CAS operation on a single value.
	// Pass previous = nil to create a new key.
	AtomicPut(key string, value []byte, previous *KVPair, options *WriteOptions) (bool, *KVPair, error)

	// Atomic delete of a single value
	AtomicDelete(key string, previous *KVPair) (bool, error)

	// Close the store connection
	Close()
}

// KVPair represents {Key, Value, Lastindex} tuple
type KVPair struct {
	Key       string
	Value     []byte
	LastIndex uint64
}

// WriteOptions contains optional request parameters
type WriteOptions struct {
	IsDir bool
	TTL   time.Duration
}

// LockOptions contains optional request parameters
type LockOptions struct {
	Value     []byte        // Optional, value to associate with the lock
	TTL       time.Duration // Optional, expiration ttl associated with the lock
	RenewLock chan struct{} // Optional, chan used to control and stop the session ttl renewal for the lock
}

// Locker provides locking mechanism on top of the store.
// Similar to `sync.Lock` except it may return errors.
type Locker interface {
	Lock(stopChan chan struct{}) (<-chan struct{}, error)
	Unlock() error
}


// Put tree
// eg bucketa/bucketb/bucket_c/key => value
// Get Tree