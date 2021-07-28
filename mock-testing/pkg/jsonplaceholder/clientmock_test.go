package jsonplaceholder_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/rnkoaa/go-labs/mock-testing/pkg/jsonplaceholder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var client jsonplaceholder.Client

// mocks the post service
type mockPostClient struct {
	mock.Mock
}

func (p mockPostClient) Get(ctx context.Context, id int) (jsonplaceholder.Post, error) {

	return jsonplaceholder.Post{}, nil
}

func (p mockPostClient) Create(ctx context.Context, post jsonplaceholder.Post) (bool, error) {
	args := p.Called(ctx, post)
	return args.Bool(0), args.Error(1)
}

func (p mockPostClient) Update(ctx context.Context, post jsonplaceholder.Post) (bool, error) {

	return false, nil
}

func (p mockPostClient) Delete(ctx context.Context, id int) (bool, error) {

	return false, nil
}

func (p mockPostClient) List(ctx context.Context) ([]jsonplaceholder.Post, error) {

	return nil, nil
}

func TestPostNotCreated(t *testing.T) {
	client = jsonplaceholder.NewClient()
	m := &mockPostClient{}
	client.Post = m
	ctx := context.Background()
	post := jsonplaceholder.Post{}

	// mock the call to return true - call is assumed to be successful
	m.On("Create", ctx, post).Return(false, fmt.Errorf("Post cannot be created"))

	ok, err := client.Post.Create(context.Background(), jsonplaceholder.Post{})

	assert.NotNil(t, err)
	assert.False(t, ok)
}
func TestClientCreated(t *testing.T) {
	client = jsonplaceholder.NewClient()
	m := &mockPostClient{}
	client.Post = m
	ctx := context.Background()
	post := jsonplaceholder.Post{}

	// mock the call to return true - call is assumed to be successful
	m.On("Create", ctx, post).Return(true, nil)

	ok, err := client.Post.Create(context.Background(), jsonplaceholder.Post{})

	assert.Nil(t, err)
	assert.True(t, ok)
}
