package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// User -
type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Address  *Address `json:"address,omitempty"`
}

// ToString -
func (u *User) ToString() string {
	return fmt.Sprintf("User{Id: %d, Name: %s, Username: %s, Email: %s}", u.ID, u.Name, u.Username, u.Email)
}

// Address -
type Address struct {
	Street      string            `json:"address"`
	Suite       string            `json:"suite,omitempty"`
	City        string            `json:"city,omitempty"`
	ZipCode     string            `json:"zipcode,omitempty"`
	GeoLocation map[string]string `json:"geo,omitempty"`
}

func fetchAllUsersSync(size int) *Result {
	users := make([]*User, 0, size)
	for id := 1; id <= size; id++ {
		url := fmt.Sprintf("%s/%d", usersURLBase, id)
		body, err := fetch(url)
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Printf("error unmarshalling user in fetchAllUsersSync: %v", err.Error())
			// return &Result{
			// 	Error:   err,
			// 	User: nil,
			// }
		}
		users = append(users, &user)
	}

	return &Result{
		Error: nil,
		Users: users,
	}
}

func fetchUsers(size, concurrency int, f func(int) *Result) *Result {
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

	responses := make([]*User, 0, size)
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
			responses = append(responses, r.User)
		}
	}

	// finally return the results
	return &Result{
		Error: nil,
		Users: responses,
	}
}

func fetchAllUsersOneGoRoutine(size int) <-chan *Result {
	out := make(chan *Result)
	go func() {
		out <- fetchAllUsersSync(size)
		close(out)
	}()
	return out
}

func fetchSingleUserSync(id int) *Result {
	url := fmt.Sprintf("%s/%d", usersURLBase, id)
	body, err := fetch(url)
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return &Result{
			Error: err,
			User:  nil,
		}
	}
	return &Result{
		Error: nil,
		User:  &user,
	}
}
