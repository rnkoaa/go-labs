package main

import "os"

func main() {
	dir := "."
	if len(os.Args) >= 2 {
		dir = os.Args[1]
	}
	digests(dir)
}
