package main

import (
	"context"
	"fmt"

	"github.com/rnkoaa/go-labs/json-placeholder/pkg/jsonplaceholder"
)

func main() {
	ctx := context.Background()
	// ctx = context.WithTimeout(ctx, 5*time.Second)
	// httpClient := http.DefaultClient
	client := jsonplaceholder.NewClient(nil)
	// client.SetBaseURL(t)
	post, _, err := client.Post.Get(ctx, 2)
	if err != nil {
		fmt.Printf("error requesting post: %v\n", err)
	}
	fmt.Println(post)
}
