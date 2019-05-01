package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// Photo -
type Photo struct {
	AlbumID      int    `json:"albumId,omitemtpy"`
	ID           int    `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	URL          string `json:"url,omitempty"`
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`
}

// ToString -
func (p *Photo) ToString() string {
	return fmt.Sprintf("Photo{Id: %d, UserId: %d, Title: %s, URL: %s}", p.ID, p.AlbumID, p.Title, p.URL)
}

func fetchAllPhotosSync(size int) *Result {
	photos := make([]*Photo, 0, size)
	for id := 1; id <= size; id++ {
		url := fmt.Sprintf("%s/%d", photosURLBase, id)
		body, err := fetch(url)
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		var photo Photo
		err = json.Unmarshal(body, &photo)
		if err != nil {
			log.Printf("error unmarshalling photo in fetchAllPhotosSync: %v", err.Error())
		}
		photos = append(photos, &photo)
	}

	return &Result{
		Error:  nil,
		Photos: photos,
	}
}

func fetchPhotos(size, concurrency int, f func(int) *Result) *Result {
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

	responses := make([]*Photo, 0, size)
	errCount := 0
	for r := range results {
		if r.Error != nil {
			log.Printf("error :%v", r.Error.Error())
			errCount++
			if errCount >= 3 {
				// fmt.Println("Too many errors, breaking!")
				log.Printf("Too many errors, breaking.")
				break
			}
		} else {
			responses = append(responses, r.Photo)
		}
	}

	// finally return the results
	return &Result{
		Error:  nil,
		Photos: responses,
	}
}

func fetchSinglePhotoSync(id int) *Result {
	url := fmt.Sprintf("%s/%d", photosURLBase, id)
	body, err := fetch(url)
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	var photo Photo
	err = json.Unmarshal(body, &photo)
	if err != nil {
		return &Result{
			Error: err,
			Photo: nil,
		}
	}
	return &Result{
		Error: nil,
		Photo: &photo,
	}
}
