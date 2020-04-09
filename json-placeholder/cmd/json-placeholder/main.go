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
	post, _, err := client.Post.Get(ctx, 2)
	if err != nil {
		fmt.Printf("error requesting post: %v\n", err)
	}
	fmt.Println(post)
}
