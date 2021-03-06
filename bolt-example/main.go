package main

// https://gist.github.com/joyrexus/22c3ef0984ed957f54b9
// https://play.golang.org/p/P5df-L-Jco
// https://medium.com/wtf-dial/wtf-dial-boltdb-a62af02b8955

// this is what i want
// https://sourcegraph.com/github.com/docker/libkv/-/blob/store/boltdb/boltdb.go
// https://github.com/docker/libkv/blob/master/store/boltdb/boltdb.go
import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var worldBucket = []byte("world")

func main() {
	db, err := bolt.Open("bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	key := "hello"
	value := "Hello World!"

	// store some data
	err = put(db, worldBucket, []byte(key), []byte(value))

	if err != nil {
		log.Fatal(err)
	}

	// retrieve the data
	val, err := get(db, worldBucket, []byte(key))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("value retrieved :%s", val)
	log.Printf(">>>>>>>>>>>>>>")
	list, err := list(db, worldBucket)

	if err != nil {
		log.Fatal(err)
	}
	for k, v := range list {
		log.Printf("Key: %s => value: %s", k, v)
	}
}

func put(db *bolt.DB, bucket, key, value []byte) error {
	// store some data
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
}

func get(db *bolt.DB, bucket, key []byte) (string, error) {
	var s []byte
	// retrieve the data
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", worldBucket)
		}

		v := bucket.Get(key)
		s = make([]byte, len(v))
		copy(s, v)

		// fmt.Println(string(val))

		return nil
	})
	return string(s), err
}

func list(db *bolt.DB, bucket []byte) (map[string]string, error) {
	results := make(map[string]string, 0)
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			results[string(k)] = string(v)
		}
		return nil
	})
	return results, err
}
