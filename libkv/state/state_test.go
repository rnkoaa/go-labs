package state

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"testing"
)

type Person struct {
	FirstName string
	Age       int
}

func (p *Person) Bytes() []byte {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	err := e.Encode(p)
	if err != nil {
		log.Printf("error encoding to bytes")
		return nil
	}
	return b.Bytes()
}

func TestInit(t *testing.T) {
	Init()
	// t.Fail()

	p := &Person{
		FirstName: "Richard",
		Age:       36,
	}

	Save("person-1", p)

	_, err := os.Stat("./app.state")
	if err != nil && os.IsNotExist(err) {
		t.Fatal(err)
		t.Fail()
	}

	b, found := Load("person-1")
	if !found {
		t.Fatalf("should have found a value with key: %s", "person-1")
		t.Fail()
	}

	if b == nil {
		t.Fatal("value returned is not valid. expect a non nil, got a nil")
		t.Fail()
	}

	// if _, err := os.Stat("./app.state.db"); os.IsNotExist(err) {
	// 	log.Printf("path exists")
	// 	// path/to/whatever does not exist
	// } else if os.IsNotExist(err) {
	// 	// path/to/whatever does *not* exist
	// 	// t.Fatalf("failed to to initialize ")
	// 	t.Fail()
	// } else {

	// 	// t.Log("Error: %v", err)
	// 	// Schrodinger: file may or may not exist. See err for details.

	// 	// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

	// }
	// Close()
}
