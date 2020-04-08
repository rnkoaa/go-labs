package jsonplaceholder

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// baseURL := "https://jsonplaceholder.typicode.com/posts"
	url := fmt.Sprintf("%s/posts/%d", p.client.BaseURL, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Post{}, nil, err
	}
	req = req.WithContext(ctx)
	res, err := p.client.client.Do(req)
	if err != nil {
		return Post{}, nil, err
	}

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		b, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return Post{}, nil, err
		}
		post, err := deserialize(b)
		return post, &Response{res}, err
	}

	return Post{}, &Response{res}, nil

}
