package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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
	fmt.Println("Hello, Gobs;")

	richard := NewPerson("John", "Smith", 36)
	sha := richard.checkSum()

	fmt.Printf("1. Sha 256 :%x\n", sha)

	richard2 := NewPerson("John", "Smith", 36)
	sha2 := richard2.checkSum()

	fmt.Printf("2. Sha 256: %x\n", sha2)
	if sha == sha2 {
		fmt.Println("Both objects are equal")
	} else {
		fmt.Println("They are not equal")
	}

	// var b bytes.Buffer
	// e := gob.NewEncoder(&b)
	// if err := e.Encode(richard); err != nil {
	// log.Fatalln("Error encoding richard.")
	// }

	// by := b.Bytes()

	// fmt.Println(by)
	// sha256 := sha256.Sum256(by)
	// log.Printf("sha256 checksum: %x", sha256)
	//	log.Println("Encoded struct: ", b)
	//	var richardDecode Person
	//	d := gob.NewDecoder(&b)
	//	if err := d.Decode(&richardDecode); err != nil {
	//		log.Fatalln("error decoding bytes")
	//	}

	//	log.Printf("Decoded Richard :%s\n", richardDecode.ToString())
	//	fmt.Printf("Person :%#v\n", richard.ToString())
}
