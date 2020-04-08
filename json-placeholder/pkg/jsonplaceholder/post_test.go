package jsonplaceholder

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetJsonPlaceHolder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s\n", r.URL)
		// if r.URL == "/posts/2" {

		// }
		fmt.Fprintln(w, `{"id": 2, "title": "hello, World", "body": "hello, world body", "userId": 1}`)
	}))
	defer ts.Close()

	testClient := ts.Client()
	if testClient == nil {
		t.Fatalf("test client should not be nil")
	}
	client := NewClient(testClient)
	client.SetBaseURL(ts.URL)
	ctx := context.TODO()

	res, _, err := client.Post.Get(ctx, 2)
	if err != nil {
		t.Fatalf("error while requesting post, object, %v", err)
	}

	if &res == nil {
		t.Fatal("error while expecting post object.")
	}

	expectedRes := Post{
		ID:     2,
		Title:  "hello, World",
		Body:   "hello, world body",
		UserID: 1,
	}

	if !reflect.DeepEqual(res, expectedRes) {
		t.Fatalf("expected response does not match")
	}
}
