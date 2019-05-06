package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	//  "github.com/doublerebel/bellows"
	"github.com/doublerebel/bellows"
	"github.com/ghodss/yaml"
	"github.com/hashicorp/vault/api"
	// yaml "gopkg.in/yaml.v2"
)

// SecretData -
type SecretData struct {
	originalKey string
	vaultKey    string
	value       string
}

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

// github.com/jeremywohl/flatten
func readFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Flatten takes a map and returns a new one where nested maps are replaced
// by dot-delimited keys.
func flattenMap(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := flattenMap(child)
			for nk, nv := range nm {
				o[k+"."+nk] = nv
			}
		default:
			o[k] = v
		}
	}
	return o
}

func processSecretFile(c *api.Client, filename string) {
	secret, err := readSecretFile("secret.yml")
	if err != nil {
		log.Printf("error reading secret file: %s", err.Error())
	}
	maps := parseSecret(secret)
	f := flattenSecret(maps)

	secretData := make([]*SecretData, 0)
	for k, v := range f {
		// fmt.Printf("Key : %s => value : %s\n", k, v.(string))
		secretData = append(secretData, &SecretData{originalKey: k, vaultKey: v.(string)})
	}

	errs := loadSecrets(c, secretData)

	if len(errs) > 0 {
		log.Printf("there are %d errors while loading secrets", len(errs))
		for k, e := range errs {
			log.Printf("%s had error %#v", k, e)
		}
	} else {
		log.Printf("There are no errors while loading secrets. proceeding")
		// for _, v := range secretData {
		// 	fmt.Printf("%#v\n", v)
		// }

		createSecretFile(secretData)
	}
}

func createSecretFile(secretData []*SecretData) {
	secrets := make(map[string]interface{}, len(secretData))
	for _, s := range secretData {
		secrets[s.originalKey] = s.value
	}

	// put them back into its original yaml structure
	y := bellows.Expand(secrets)
	ymlData, err := yaml.Marshal(y)
	if err != nil {
		log.Printf("error converting data into yml")
	}

	// this is the yamlData to post
	// calculate the sha256 of this file and update.
	log.Printf("%s", string(ymlData))
}

func loadSecrets(c *api.Client, secretData []*SecretData) map[string]error {
	errs := make(map[string]error, 0)
	errCount := 0
	for _, s := range secretData {
		vaultKey := s.vaultKey
		value, err := ReadSecret(c, vaultKey)
		if err != nil {
			errs[s.originalKey] = err
			errCount++
			if errCount >= 3 {
				break
			}
		}
		s.value = value
	}
	return errs
}
