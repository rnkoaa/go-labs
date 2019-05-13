package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// CreateNestedBucketsNew - function to create
// nested buckets from an array of Strings - my implementation
func CreateNestedBucketsNew(buckets []string) (err error) {
	db, dberr := bolt.Open(dbname, dbperms, options)
	if dberr != nil {
		log.Fatal(dberr)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) (err error) {
		var bkt *bolt.Bucket

		for index, bucket := range buckets {
			if index == 0 {
				bkt, err = tx.CreateBucketIfNotExists([]byte(bucket))
			} else {
				bkt, err = bkt.CreateBucketIfNotExists([]byte(bucket))
			}

			if err != nil {
				return fmt.Errorf("Error creating nested bucket [%s]: %v", bucket, err)
			}
		}
		return err
	})
	return err
}
