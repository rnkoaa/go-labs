package jsonplaceholder

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// PostService interface exposes all the available methods needed to interact with posts
// on jsonplaceholder.com
type PostService interface {
	Get(ctx context.Context, id int) (Post, *Response, error)
}

// PostServiceClient implements the PostService interface
type PostServiceClient struct {
	client *Client
}

var _ PostService = &PostServiceClient{}

// Post represents a Post response from json placeholder
type Post struct {
	Body   string `json:"body"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	UserID int    `json:"userId"`
}

func deserialize(response []byte) (Post, error) {
	var p Post
	err := json.Unmarshal(response, &p)
	return p, err
}

// Get post from json placeholder given the id of the post
func (p *PostServiceClient) Get(ctx context.Context, id int) (Post, *Response, error) {
	path := fmt.Sprintf("/posts/%d", id)
	req, err := p.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return Post{}, nil, err
	}

	var post Post
	res, err := p.client.Do(ctx, req, &post)
	if err != nil {
		return Post{}, nil, err
	}
	return post, res, nil
}
