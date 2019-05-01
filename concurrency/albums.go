package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// Album -
type Album struct {
	UserID int    `json:"userId,omitempty"`
	ID     int    `json:"id"`
	Title  string `json:"title,omitempty"`
}

// ToString -
func (a *Album) ToString() string {
	return fmt.Sprintf("Album{Id: %d, UserId: %d, Title: %s}", a.ID, a.UserID, a.Title)
}

func fetchAllAlbumsSync(size int) *Result {
	albums := make([]*Album, 0, size)
	for id := 1; id <= size; id++ {
		url := fmt.Sprintf("%s/%d", albumsURLBase, id)
		body, err := fetch(url)
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		var album Album
		err = json.Unmarshal(body, &album)
		if err != nil {
			log.Printf("error unmarshalling album in fetchAllAlbumsSync: %v", err.Error())
			// return &Result{
			// 	Error:   err,
			// 	Album: nil,
			// }
		}
		albums = append(albums, &album)
	}

	return &Result{
		Error:  nil,
		Albums: albums,
	}
}

func fetchAlbums(size, concurrency int, f func(int) *Result) *Result {
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

	responses := make([]*Album, 0, size)
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
			responses = append(responses, r.Album)
		}
	}

	// finally return the results
	return &Result{
		Error:  nil,
		Albums: responses,
	}
}

func fetchSingleAlbumSync(id int) *Result {
	url := fmt.Sprintf("%s/%d", albumsURLBase, id)
	body, err := fetch(url)
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	var album Album
	err = json.Unmarshal(body, &album)
	if err != nil {
		return &Result{
			Error: err,
			Album: nil,
		}
	}
	return &Result{
		Error: nil,
		Album: &album,
	}
}
