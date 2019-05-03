package main

import (
	"fmt"
	"log"
)

func main() {
	secret, err := readSecretFile("secret.yml")
	if err != nil {
		log.Printf("error reading secret file: %s", err.Error())
	}
	maps := parseSecret(secret)
	f := flattenSecret(maps)

	for k, v := range f {
		fmt.Printf("Key : %s => value : %s\n", k, v.(string))
	}
}
