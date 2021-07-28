package jsonplaceholder

import "context"

type Post struct {
	ID      int    `json:"id"`
	UsertID int    `json:"userId"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type PostService interface {
	Get(context.Context, int) (Post, error)
	List(context.Context) ([]Post, error)
	Create(context.Context, Post) (bool, error)
	Update(context.Context, Post) (bool, error)
	Delete(context.Context, int) (bool, error)
}

type PostServiceClient struct{}

var _ PostService = &PostServiceClient{}

func (p PostServiceClient) Get(ctx context.Context, id int) (Post, error) {

	return Post{}, nil
}

func (p PostServiceClient) Create(ctx context.Context, post Post) (bool, error) {

	return false, nil
}

func (p PostServiceClient) Update(ctx context.Context, post Post) (bool, error) {

	return false, nil
}

func (p PostServiceClient) Delete(ctx context.Context, id int) (bool, error) {

	return false, nil
}

func (p PostServiceClient) List(ctx context.Context) ([]Post, error) {

	return nil, nil
}
