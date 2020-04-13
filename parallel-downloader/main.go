// https://gist.github.com/fracasula/b579d52daf15426e58aa133d0340ccb0
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
}

func main() {
	// var wg sync.WaitGroup
	// defer wg.Wait()
	// jsonResponses := downloadItems(wg)
	// jsonResponses := make(chan string)
	// wg.Add(len(urls))

	for _, url := range urls {
		res := processUrl(url)
		if res.err != nil {
			fmt.Printf("url [%s] error %v\n", url, res.err)
		} else {
			fmt.Printf("url: %s, response %s\n", url, res.res)
		}
	}

	// go func() {
	// 	for response := range jsonResponses {
	// 		fmt.Println(response)
	// 	}
	// }()
}

type result struct {
	err error
	res string
}

func processUrls(urls []string) {
	numFinders := 4
	// finders := make([]<-chan int, numFinders)
	for i := 0; i < numFinders; i++ {
		// finders[i] = primeFinder(done, randIntStream)
	}
}

func processUrl(url string) result {
	res, err := http.Get(url)
	if err != nil {
		return result{
			err: err,
			res: "",
		}
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result{
			err: errors.New("error reading response body"),
			res: "",
		}
	}
	return result{
		err: nil,
		res: string(body),
	}
}
