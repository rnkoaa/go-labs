package main

// GetSecretKeys -
func GetSecretKeys(secrets map[string]interface{}) []string {
	keys := make([]string, 0)

	for key := range secrets {
		keys = append(keys, key)
	}
	return keys
}
