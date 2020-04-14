package main

import (
	"encoding/json"
	domain "github.com/rnkoaa/go-labs/json-server/pkg/domain"
	"io/ioutil"
	"log"
)

// import "../../pkg/domain"

type InMemoryDatabase struct {
	// Posts []*Post `json:"posts"`
}

// // Database -
type Database struct {
	Albums   []*domain.Album   `json:"albums"`
	Comments []*domain.Comment `json:"comments"`
	Photos   []*domain.Photo   `json:"photos"`
	Posts    []*domain.Post    `json:"posts"`
	Users    []*domain.User    `json:"users"`
	Todos    []*domain.Todo    `json:"todos"`
}

func NewDatabase(path string) *Database {
	dbPath := "db.json"
	if path != "" {
		dbPath = path
	}

	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		log.Fatalf("error reading database file %v", err)
	}

	var d Database
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Fatalf("error unmarshalling json file %v", err)
	}
	return &d
}

