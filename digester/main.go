package main

import (
	"os"
)

func main() {
	loc := os.Args[1]
	// serial(loc)
	//	parallelRun(loc)
	bounded(loc)

}
