package jsonplaceholder

import (
	"context"
	"fmt"
	"net/http"
)

/*
{
      "userId": 1,
      "id": 20,
      "title": "ullam nobis libero sapiente ad optio sint",
      "completed": true
	}
*/
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

/*
{
  "id": 1,
  "name": "Leanne Graham",
  "username": "Bret",
  "email": "Sincere@april.biz",
  "address": {
    "street": "Kulas Light",
    "suite": "Apt. 556",
    "city": "Gwenborough",
    "zipcode": "92998-3874",
    "geo": {
      "lat": "-37.3159",
      "lng": "81.1496"
    }
  },
  "phone": "1-770-736-8031 x56442",
  "website": "hildegard.org",
  "company": {
    "name": "Romaguera-Crona",
    "catchPhrase": "Multi-layered client-server neural-net",
    "bs": "harness real-time e-markets"
  }
}
*/

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
