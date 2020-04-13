// https://gist.github.com/fracasula/b579d52daf15426e58aa133d0340ccb0
// https://blog.afoolishmanifesto.com/posts/golang-concurrency-patterns/
package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var urls = []string{
	"https://jsonplaceholder.typicode.com/posts/1",
	"https://jsonplaceholder.typicode.com/posts/2",
	"https://jsonplaceholder.typicode.com/posts/3",
	"https://jsonplaceholder.typicode.com/posts/4",
	"https://jsonplaceholder.typicode.com/posts/5",
	"https://jsonplaceholder.typicode.com/posts/6",
	"https://jsonplaceholder.typicode.com/posts/7",
	"https://jsonplaceholder.typicode.com/posts/8",
	"https://jsonplaceholder.typicode.com/posts/9",
	"https://jsonplaceholder.typicode.com/posts/10",
	"https://jsonplaceholder.typicode.com/posts/11",
	"https://jsonplaceholder.typicode.com/posts/12",
}

type result struct {
	err error
	res string
}

func request(ctx context.Context, wg *sync.WaitGroup, url string, ch chan result) {
	res, err := http.Get(url)
	if err != nil {
		ch <- result{
			err: err,
			res: "",
		}
		wg.Done()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ch <- result{
			err: errors.New("error reading response body"),
			res: "",
		}
		wg.Done()
	}
	res.Body.Close()
	ch <- result{
		err: nil,
		res: string(body),
	}
	wg.Done()
}

func makeRequests(ctx context.Context, wg *sync.WaitGroup, urls []string, ch chan result) {
	wg.Add(len(urls))
	for _, u := range urls {
		request(ctx, wg, u, ch)
	}
	close(ch)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	ch := make(chan result)

	go makeRequests(ctx, &wg, urls, ch)

	for res := range ch {
		if res.err != nil {
			fmt.Println(res.err)
		} else {
			fmt.Println(res.res)
		}
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// block until either an error or OS-level signals
	// to shutdown gracefully
	select {
	case <-sigChan:
		fmt.Printf("Shutdown signal received... closing server")
		cancel()
	}
	wg.Wait()
}
