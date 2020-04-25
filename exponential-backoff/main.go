package main

import (
	"fmt"
	"math"
	"time"
)

var (
	baseSleepTime = 100 * time.Millisecond
)

func main() {
	fmt.Println("vim-go")
	for i := 0; i < 15; i++ {
		fmt.Printf("Sleep Time: %v\n", getSleepTime(i))
	}
}

func getSleepTime(count int) time.Duration {
	intermediate := math.Pow(2, float64(count))
	sleepTime := int(intermediate) * 100
	return time.Duration(sleepTime) * time.Millisecond
}
