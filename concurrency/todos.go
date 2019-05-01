package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// Todo -
type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// ToString -
func (t *Todo) ToString() string {
	return fmt.Sprintf("Todo{Id: %d, UserId: %d, Title: %s, Completed: %v}",
		t.ID, t.UserID, t.Title, t.Completed)
}

func fetchAllTodosSync(size int) *Result {
	todos := make([]*Todo, 0, size)
	for id := 1; id <= size; id++ {
		url := fmt.Sprintf("%s/%d", todoURLBase, size)
		body, err := fetch(url)
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		var post Todo
		err = json.Unmarshal(body, &post)
		if err != nil {
			log.Printf("error unmarshalling post in fetchAllTodosSync: %v", err.Error())
		}
		todos = append(todos, &post)
	}

	return &Result{
		Error: nil,
		Todos: todos,
	}
}

func fetchTodos(size, concurrency int, f func(int) *Result) *Result {
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

	responses := make([]*Todo, 0, size)
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
			responses = append(responses, r.Todo)
		}
	}

	// finally return the results
	return &Result{
		Error: nil,
		Todos: responses,
	}
}

func fetchAllTodosOneGoRoutine(size int) <-chan *Result {
	out := make(chan *Result)
	go func() {
		out <- fetchAllTodosSync(size)
		close(out)
	}()
	return out
}

func fetchSingleTodoSync(id int) *Result {
	url := fmt.Sprintf("%s/%d", todoURLBase, id)
	body, err := fetch(url)
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	var post Todo
	err = json.Unmarshal(body, &post)
	if err != nil {
		return &Result{
			Error: err,
			Todo:  nil,
		}
	}
	return &Result{
		Error: nil,
		Todo:  &post,
	}
}
