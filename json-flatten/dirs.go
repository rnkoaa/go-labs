package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/doublerebel/bellows"
	// "gopkg.in/yaml.v2"
	"github.com/ghodss/yaml"
)

func readSecretFile(fileName string) ([]byte, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("unable to read secret file")
	}
	return b, err
}

func parseSecret(content []byte) []byte {
	// m := make(map[interface{}]interface{})
	// err := yaml.Unmarshal(content, m)
	// if err != nil {
	// 	log.Printf("error unmarshalling yaml to interface %#v", err.Error())
	// 	return nil
	// }
	j2, err := yaml.YAMLToJSON(content)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}
	return j2
	// return m
}

func flattenSecret(content []byte) map[string]interface{} {
	var b map[string]interface{}
	e := json.Unmarshal(content, &b)
	if e != nil {
		log.Printf("error unmarshalling json to bytes, %#v", e.Error())
	}
	flattened := bellows.Flatten(b)
	// b, err := json.Marshal(flattened)
	// if err != nil {
	// 	log.Printf("error marshalling content into json", err.Error())
	// }

	return flattened
}
