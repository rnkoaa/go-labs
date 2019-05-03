package main

import (
	//	"fmt"

	//	"github.com/hashicorp/vault/api"
	"log"
)

// var token = os.Getenv("TOKEN")
var token = "f1d8c4ab-2a04-03e0-8530-5afe0ab13f4f"
var vaultAddress = "http://localhost:8200"

// var vaultAddress = os.Getenv("VAULT_ADDR")

func main() {
	/*
		config := &api.Config{
			Address: vaultAddress,
		}
		client, err := api.NewClient(config)
		if err != nil {
			fmt.Println(err)
			return
		}
		client.SetToken(token)
		c := client.Logical()
		secret, err := c.Read("secret/prod")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(secret.Data["user"])

		secretData := map[string]interface{}{
			"value": "world",
			"foo":   "bar",
			"age":   "-1",
		}
		_, err = client.Logical().Write("secret/prod/hello", secretData)
	*/
	b, err := flattenYamlFile("config/secret.yml")
	if err != nil {
		log.Fatalf("error reading file %v", err)
	}
	// log.Printf(string(b))
	for k := range b {
		log.Printf("Key :%s", k)
	}
}

/**
package main

import (
	"encoding/json"
	"fmt"
)

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
