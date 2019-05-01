package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// Comment -
type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// ToString -
func (c *Comment) ToString() string {
	return fmt.Sprintf("Comment{Id: %d, UserId: %d, Title: %s, Completed: %v}", c.ID, c.PostID, c.Name, c.Email)
}

func fetchAllCommentsSync(size int) *Result {
	comments := make([]*Comment, 0, size)
	for id := 1; id <= size; id++ {
		url := fmt.Sprintf("%s/%d", commentURLBase, id)
		body, err := fetch(url)
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		var comment Comment
		err = json.Unmarshal(body, &comment)
		if err != nil {
			log.Printf("error unmarshalling comment in fetchAllCommentsSync: %v", err.Error())
			// return &Result{
			// 	Error:   err,
			// 	Comment: nil,
			// }
		}
		comments = append(comments, &comment)
	}

	return &Result{
		Error:    nil,
		Comments: comments,
	}
}

func fetchComments(size, concurrency int, f func(int) *Result) *Result {
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

	responses := make([]*Comment, 0, size)
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
			responses = append(responses, r.Comment)
		}
	}

	// finally return the results
	return &Result{
		Error:    nil,
		Comments: responses,
	}
}

func fetchSingleCommentSync(id int) *Result {
	url := fmt.Sprintf("%s/%d", commentURLBase, id)
	body, err := fetch(url)
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	var comment Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		return &Result{
			Error:   err,
			Comment: nil,
		}
	}
	return &Result{
		Error:   nil,
		Comment: &comment,
	}
}
