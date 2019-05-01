package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// Post -
type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// ToString -
func (p *Post) ToString() string {
	return fmt.Sprintf("Post{Id: %d, UserId: %d, Title: %s}", p.ID, p.UserID, p.Title)
}

func fetchAllPostsSync(size int) *Result {
	posts := make([]*Post, 0, size)
	for id := 1; id <= size; id++ {
		url := fmt.Sprintf("%s/%d", postsURLBase, id)
		body, err := fetch(url)
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		var post Post
		err = json.Unmarshal(body, &post)
		if err != nil {
			log.Printf("error unmarshalling post in fetchAllPostsSync: %v", err.Error())
			// return &Result{
			// 	Error:   err,
			// 	Post: nil,
			// }
		}
		posts = append(posts, &post)
	}

	return &Result{
		Error: nil,
		Posts: posts,
	}
}

func fetchPosts(size, concurrency int, f func(int) *Result) *Result {
	done := make(chan struct{})
	defer close(done)
	requests := make(chan int, size)
	go func(reqSize int) {
		defer close(requests)
		for i := 1; i <= size; i++ {
			requests <- i
		}
	}(size)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	results := make(chan *Result)
	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for id := range requests {
				select {
				case <-done:
					return
				case results <- f(id):
				}
			}
		}()
	}

	responses := make([]*Post, 0, size)
	errCount := 0
	for r := range results {
		if r.Error != nil {
			log.Printf("error :%v", r.Error.Error())
			errCount++
			if errCount >= 3 {
				log.Printf("Too many errors, breaking.")
				break
			}
		} else {
			responses = append(responses, r.Post)
		}
	}

	// finally return the results
	return &Result{
		Error: nil,
		Posts: responses,
	}
}

func fetchAllPostsOneGoRoutine(size int) <-chan *Result {
	out := make(chan *Result)
	go func() {
		out <- fetchAllPostsSync(size)
		close(out)
	}()
	return out
}
func fetchSinglePostSync(id int) *Result {
	url := fmt.Sprintf("%s/%d", postsURLBase, id)
	body, err := fetch(url)
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return &Result{
			Error: err,
			Post:  nil,
		}
	}
	return &Result{
		Error: nil,
		Post:  &post,
	}
}
