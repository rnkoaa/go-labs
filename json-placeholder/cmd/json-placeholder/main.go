package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rnkoaa/go-labs/json-placeholder/pkg/jsonplaceholder"
)

const (
	baseURL        = "https://jsonplaceholder.typicode.com"
	defaultTimeout = 30 * time.Second
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	client := jsonplaceholder.NewClient(nil)
	client.SetBaseURL(baseURL)
	user, _, err := client.User.Get(ctx, 1)
	if err != nil {
		fmt.Printf("error requesting post: %v\n", err)
	}
	fmt.Println(jsonplaceholder.Stringify(user))
}
