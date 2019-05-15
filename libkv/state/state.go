package state

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"log"

	"github.com/rnkoaa/libkv"
	"github.com/rnkoaa/libkv/store"
	"github.com/rnkoaa/libkv/store/boltdb"
)

// Cached -
type Cached struct {
	hasher hash.Hash
	result [sha256.Size]byte
}

// NewCached -
func NewCached() *Cached {
	return &Cached{
		hasher: sha256.New(),
	}
}

// Sum -
func (c *Cached) Sum(data []byte) []byte {
	c.hasher.Reset()
	c.hasher.Write(data)
	return c.hasher.Sum(c.result[:0])
}

var kv store.Store
var cached *Cached

// Cacheable -
type Cacheable interface {
	Bytes() []byte
}

// Init -
func Init() {
	boltdb.Register()

	var err error
	kv, err = libkv.NewStore(
		store.BOLTDB,
		[]string{"./app.state"},
		&store.Config{Bucket: "state"},
	)

	if err != nil {
		log.Fatalln("error initialinging boltdb store ", err)
	}

	cached = NewCached()
}

// Load -
func Load(key string) ([]byte, bool) {
	kvPair, err := kv.Get(key)
	if err != nil {
		return nil, false
	}

	if kvPair == nil {
		return nil, false
	}

	v := kvPair.Value
	if v == nil {
		return nil, false
	}
	return v, true
}

// Changed -
// checks to see if the key has been changed by its checksum
func Changed(key string, v Cacheable) bool {
	b, found := Load(key)
	if !found {
		return true
	}

	str := hex.EncodeToString(b)
	checkSum := cached.Sum(b)
	if checkSum == nil {
		return false
	}
	chkStr := hex.EncodeToString(checkSum)
	return str == chkStr
}

// Save -
func Save(key string, v Cacheable) bool {
	b := v.Bytes()
	checkSum := cached.Sum(b)
	e := kv.Put(key, checkSum, &store.WriteOptions{IsDir: true})
	if e != nil {
		log.Printf("error saving %s => %v", key, e)
		return false
	}
	return true
}

// Close -
func Close() {
	kv.Close()
}

/*
// Department
	// classes

// Class
	// Teacher
	// Students
	// CourseNumber
	// MeetingTime

// Student
	// Name
	// Age
*/
