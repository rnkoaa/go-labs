package main

// https://stackoverflow.com/questions/50276670/load-a-dynamic-yaml-structure-in-go

/*
---
supplychain:
  - dist-receipts-article-watcher-prod:
      hash: 12345678892423
      resources:
        - name: 'prod/application.yml'
          hash: 12345678892423
        - name: 'prod/secret.yml'
          hash: 12345678892423
        - name: 'prod/inventory-processing-prd.jks'
          hash: 12345678892423
  - dist-receipts-article-watcher-stage:
      hash: 12345678892423
      resources:
        - name: 'prod/application.yml'
          hash: 12345678892423
        - name: 'prod/secret.yml'
          hash: 12345678892423
        - name: 'prod/inventory-processing-prd.jks'
          hash: 12345678892423
purchaseorders:
  - dist-receipts-article-watcher-prod:
      hash: 12345678892423
      resources:
        - name: 'prod/application.yml'
          hash: 12345678892423
        - name: 'prod/secret.yml'
          hash: 12345678892423
        - name: 'prod/inventory-processing-prd.jks'
          hash: 12345678892423
  - dist-receipts-article-watcher-stage:
      hash: 12345678892423
      resources:
        - name: 'prod/application.yml'
          hash: 12345678892423
        - name: 'prod/secret.yml'
          hash: 12345678892423
        - name: 'prod/inventory-processing-prd.jks'
          hash: 12345678892423
*/

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func readFile(location string) []byte {
	b, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatalf("Error reading yaml state file.")
	}
	return b
}

func main() {
	fileContent := readFile("application.state.yml")
	yamlMap := make(map[interface{}]interface{})

	err := yaml.Unmarshal(fileContent, yamlMap)
	if err != nil {
		log.Fatalf("error unmarshalling yaml content %#v", err.Error())
	}

	for key, value := range yamlMap {
		fmt.Printf("Application: %s\n", key.(string))

		clusterMap := value.(map[interface{}]interface{})
		for cluster, value := range clusterMap {
			if cluster.(string) == "hash" {
				// this is the application hash
				fmt.Printf("Application Hash: %s\n", value.(string))
			} else {

				// fmt.Printf("Cluster name: %s\n", cluster)
				clusters := value.(map[interface{}]interface{})
				for c, co := range clusters {
					if c.(string) == "hash" {
						fmt.Printf("Cluster Hash: %s\n", co.(string))
					} else {
						_ = co.([]interface{})
						// _ = co.(map[interface{}]interface{})
					}
				}
			}
		}
	}
	log.Println("Done.")
}
