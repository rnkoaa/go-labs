package jsonplaceholder

import (
	"context"
	"fmt"
	"net/http"
)

// Todo represents a todo json object response
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// User - object
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  *struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     *struct {
			Lat string `json:"lat"`
			LNG string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company *struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		BS          string `json:"bs"`
	} `json:"company"`
}

// UserService interface exposes all the available methods needed to interact with posts
// on jsonplaceholder.com
type UserService interface {
	Get(context.Context, int) (User, *Response, error)
	List(context.Context, int, int) ([]User, *Response, error)
}

// UserServiceClient implements the PostService interface
type UserServiceClient struct {
	client *Client
}

var _ UserService = &UserServiceClient{}

// List all posts in the api
func (p *UserServiceClient) List(ctx context.Context, page, size int) ([]User, *Response, error) {

	return nil, nil, nil
}

// Get User from json placeholder given the id of the post
func (p *UserServiceClient) Get(ctx context.Context, id int) (User, *Response, error) {
	path := fmt.Sprintf("/users/%d", id)
	req, err := p.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return User{}, nil, err
	}

	var user User
	res, err := p.client.Do(ctx, req, &user)
	if err != nil {
		return User{}, nil, err
	}
	return user, res, nil
}
