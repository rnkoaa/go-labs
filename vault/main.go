package main

import (
	//	"fmt"

	//	"github.com/hashicorp/vault/api"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

const (
	env  = "stage"
	team = "cost"
)

var (
	baseSecretPath = fmt.Sprintf("secret/%s/%s", team, env)
)

// var token = os.Getenv("TOKEN")
// var token = "f1d8c4ab-2a04-03e0-8530-5afe0ab13f4f"
var token = "cce8e473-9ec2-b13d-2134-5c1780ae37b3"
var vaultAddress = "http://localhost:8200"

// var vaultAddress = os.Getenv("VAULT_ADDR")

func initClient() *api.Client {
	config := &api.Config{
		Address: vaultAddress,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	client.SetToken(token)
	return client
}
func main() {
	c := initClient()
	processSecretFile(c, "secret.yml")
	// loadJSONIntoVault(c, "secrets.json")
	// dumpMap("", results)

}

// ReadSecret -
func ReadSecret(c *api.Client, secretKey string) (string, error) {
	keyPath := fmt.Sprintf("%s/%s", baseSecretPath, secretKey)
	secret, err := c.Logical().Read(keyPath)
	if err != nil {
		return "", err
	}
	secretValue := secret.Data["value"]
	if secretValue != nil {
		return secretValue.(string), nil
	}
	return "", nil
}

// ReadSecrets -
func ReadSecrets(c *api.Client, keys []string) (map[string]string, map[string]error) {
	errs := make(map[string]error, 0)
	secrets := make(map[string]string, 0)
	errCount := 0
	for _, key := range keys {
		secret, err := ReadSecret(c, key)
		if err != nil {
			errs[key] = err
			errCount++
			if errCount >= 3 {
				log.Printf("too many errors during read. exiting...")
				break
			}
		}
		secrets[key] = secret
	}
	return secrets, errs
}

// WriteSecret -
func WriteSecret(c *api.Client, key, value string) error {
	secretData := map[string]interface{}{
		"value": value,
		"age":   "-1",
	}
	keyPath := fmt.Sprintf("%s/%s", baseSecretPath, key)
	_, err := c.Logical().Write(keyPath, secretData)
	return err
}

// WriteSecrets -
func WriteSecrets(c *api.Client, secrets map[string]string) map[string]error {
	errs := make(map[string]error, 0)
	errCount := 0
	for key, value := range secrets {
		err := WriteSecret(c, key, value)
		if err != nil {
			errs[key] = err
			errCount++
			// there may be something going on, please exit and check on it.
			if errCount == 3 {
				log.Printf("too many errors, exiting...")
				break
			}
		}
	}
	return errs
}

/**
package main

import (
	"encoding/json"
	"fmt"
)



var jsonStr = `
{
  "array": [
	1,
	2,
	3
  ],
  "boolean": true,
  "null": null,
  "number": 123,
  "object": {
	"a": "b",
	"c": "d",
	"e": "f"
  },
  "string": "Hello World"
}
`

func main() {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	dumpMap("", jsonMap)
}
*/

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

func loadJSONIntoVault(c *api.Client, filename string) {
	results, err := readSecretsJSONFile("secrets.json")
	if err != nil {
		log.Fatalf("error read secrets file: %#v", err)
	}
	// for key, value := range results {
	// 	log.Printf("%s => %s", key, value)
	// }
	errs := WriteSecrets(c, results)
	if len(errs) > 0 {
		for key, e := range errs {
			log.Printf("error writing secret for %s => %#v", key, e)
		}
	}

}
func readSecretsJSONFile(filename string) (map[string]string, error) {
	b, err := readFile(filename)
	if err != nil {
		return nil, err
	}
	var jsonMap map[string]interface{}
	err = json.Unmarshal(b, &jsonMap)
	if err != nil {
		return nil, err
	}

	results := make(map[string]string, 0)
	for key, value := range jsonMap {
		results[key] = value.(string)
	}
	return results, nil
}
