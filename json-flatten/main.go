package main

import (
	"fmt"
	"log"
)

// "github.com/doublerebel/bellows"

func main() {
	secret, err := readSecretFile("secret.yml")
	if err != nil {
		log.Printf("error reading secret file: %s", err.Error())
	}
	maps := parseSecret(secret)
	// fmt.Println(string(maps))

	// b, e := json.Marshal(maps)
	// if e != nil {
	// 	log.Fatal("error marshalling maps to json => ", e)
	// }

	// log.Printf(string(b))
	f := flattenSecret(maps)
	// fmt.Println(f)

	for k, v := range f {
		fmt.Printf("Key : %s => value : %s\n", k, v.(string))
		// get the value from value for v.(string)
	}
	// s, e := json.Marshal(f)
	// if e != nil {
	// 	log.Printf("error marshalling flattened json: %#v", e.Error())
	// }

	// fmt.Printf(string(s))

	// nested := make(map[string]interface{})
	// nested["hello"] = "World"

	// >>>>>> ----- >>>>>>
	// nested2 := make(map[string]interface{})
	// nested2["name"] = "richard"
	// nested["second"] = nested2
	// flattened := bellows.Flatten(nested)
	// fmt.Printf("%#v\n", flattened)
	// fmt.Println("---------------------------------------------------")
	// fmt.Println("-------- Secret Keys -----------------------------")
	// secretKeys := GetSecretKeys(flattened)
	// fmt.Printf("%#v\n", secretKeys)
	// fmt.Println("---------------------------------------------------")

	// fmt.Println("expanded -> lists")
	// expanded := bellows.Expand(flattened)
	// fmt.Printf("%#v\n", expanded)
}
