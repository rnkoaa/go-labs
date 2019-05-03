package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/doublerebel/bellows"
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
	b, err := yaml.YAMLToJSON(content)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}
	return b
}

func flattenSecret(content []byte) map[string]interface{} {
	var b map[string]interface{}
	e := json.Unmarshal(content, &b)
	if e != nil {
		log.Printf("error unmarshalling json to bytes, %#v", e.Error())
	}
	return bellows.Flatten(b)
}
