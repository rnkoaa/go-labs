package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// Person -
type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

// ToString -
func (p *Person) ToString() string {
	return fmt.Sprintf("{Name: %s %s, Age: %d}", p.Firstname, p.Lastname, p.Age)
}

// NewPerson -
func NewPerson(firstName, lastName string, age int) *Person {
	return &Person{
		Firstname: firstName,
		Lastname:  lastName,
		Age:       age,
	}
}

func (p *Person) Bytes() ([]byte, error) {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	err := e.Encode(p)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (p *Person) checkSum() [32]byte {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(p); err != nil {
		log.Fatalln("error encoding person.")

	}

	by := b.Bytes()
	sha256 := sha256.Sum256(by)
	return sha256
}

func main() {
	// fmt.Println("Hello, Gobs;")

	richard := NewPerson("John", "Smith", 36)
	// sha := richard.checkSum()

	// fmt.Printf("1. Sha 256 :%x\n", sha)

	// richard2 := NewPerson("John", "Smith", 36)
	// sha2 := richard2.checkSum()

	// fmt.Printf("2. Sha 256: %x\n", sha2)
	// if sha == sha2 {
	// 	fmt.Println("Both objects are equal")
	// } else {
	// 	fmt.Println("They are not equal")
	// }
	cached := NewCached()

	by, err := richard.Bytes()
	if err != nil {
		log.Fatalf("error generating bytes: %v", err)
	}
	sum := cached.Sum(by)
	s := fmt.Sprintf("%x", sum)
	en := hex.EncodeToString(sum)
	log.Printf("checksum: %s\nHex: %s", s, en)
}
