package main

import (
	"io/ioutil"
	"log"
	bellows "github.com/doublerebel/bellows"
	yaml "gopkg.in/yaml.v2"
)

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

func flattenYamlFile(filename string) (map[string]interface{}, error) {
	b, err := readFile(filename)
	if err != nil {
		return nil, err
	}
	yamlRaw := make(map[interface{}]interface{})

	err = yaml.Unmarshal(b, yamlRaw)
	if err != nil {
		return nil, err
	}

	log.Printf("%#v", yamlRaw)

	flat := bellows.Flatten(yamlRaw)
	log.Printf("%#v", flat)


// yamlMap := make(map[string]interface{}, len(yamlRaw))
// 	for k, v := range flat {
// 		yamlMap[k] = v 
// 	}
	return flat, nil
}
