package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rnkoaa/go-labs/json-placeholder/pkg/jsonplaceholder"
)

func main() {
	ctx := context.Background()
	// ctx = context.WithTimeout(ctx, 5*time.Second)
	client := jsonplaceholder.NewClient(http.DefaultClient)
	post, _, err := client.Post.Get(ctx, 2)
	if err != nil {
		fmt.Printf("error requesting post: %v\n", err)
	}
	fmt.Println(post)
}
