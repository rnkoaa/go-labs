package main

import (
	"fmt"
	"github.com/rnkoaa/libkv"
	"github.com/rnkoaa/libkv/store"
	"github.com/rnkoaa/libkv/store/boltdb"
	"log"
)

func main() {
	boltdb.Register()

	kv, err := libkv.NewStore(
		store.BOLTDB,
		[]string{"./__boltdbtest"},
		&store.Config{Bucket: "boltDBTest"},
	)

	if err != nil {
		log.Fatalln("error initialinging boltdb store ", err)
	}

	if kv != nil {
		//kv.Put("mKey", []byte("mValue"), &store.WriteOptions{})
		_ = kv.Put("mKey", []byte("mValue"), &store.WriteOptions{IsDir: true})
		_ = kv.Put("mKey/keybase", []byte("mValue-01"), &store.WriteOptions{IsDir: true})
		_ = kv.Put("m-1/keybase", []byte("mValue-01"), &store.WriteOptions{IsDir: true})
	}

	kvPair, err := kv.Get("mKey")
	if err != nil {
		log.Printf("error retrieving kv pair: %v", err)
	}

	if kvPair != nil {
		log.Printf("kvPair %v", string(kvPair.Value))
	}

	entries, err := kv.List("mKey")
	for _, pair := range entries {
		fmt.Printf("key=%v - value=%v\n", pair.Key, string(pair.Value))
	}

}
