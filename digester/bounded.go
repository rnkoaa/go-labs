package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

var concurrency = 2

func bounded(loc string) {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	m, err := boundedDigester(loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}

}

func boundedDigester(loc string) (map[string][sha256.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles(done, loc)

	// start a fixed number of goroutines to read and digest files.
	c := make(chan result)
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			digester(done, paths, c) //HLc
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()
	// End of pipeline

	m := make(map[string][sha256.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	// check whether the walk failed.
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}

func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{path, sha256.Sum256(data), err}:
		case <-done:
			return
		}
	}
}

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		// close the paths channel after walk returns.
		defer close(paths)

		//No select needed for this send, since errc is buffered
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("Walk cancelled")
			}
			return nil
		})
	}()
	return paths, errc
}
