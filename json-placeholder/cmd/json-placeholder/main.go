package main

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/rnkoaa/go-labs/json-placeholder/pkg/jsonplaceholder"
)

const (
	baseURL         = "https://jsonplaceholder.typicode.com"
	defaultTimeout  = 30 * time.Second
	baseSleepFactor = 200
	maxRetries      = 6
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	client := jsonplaceholder.NewClient(nil)
	client.SetBaseURL(baseURL)
	runMultipleTimes(ctx, client, 1)
}

func getSleepTime(count, sleepFactor int) time.Duration {
	intermediate := math.Pow(2, float64(count))
	sleepTime := int(intermediate) * sleepFactor
	return time.Duration(sleepTime) * time.Millisecond
}

func backOffRequest(count, sleepFactor int) {
	sleepTime := getSleepTime(count, sleepFactor)
	time.Sleep(sleepTime)
}

func runMultipleTimes(ctx context.Context, client *jsonplaceholder.Client, userId int) {
	totalCount := 10
	retries := 0
	count := 0
	for {
		fmt.Printf("Count -> %d\n", count)
		if count == totalCount {
			fmt.Println("We have reached max count, returning")
			break
		}
		retries++
		count++
		user, _, err := client.User.Get(ctx, 1)
		if err != nil {
			// check error type if it is a retryable error, do not break, let it retry
			// SUCCESS,
			// NOT_READY,
			// THROTTLED,
			/*if (Results.SUCCESS == result) {
			      retry = false;
			  } else if (Results.NOT_READY == result) {
			      retry = true;
			  } else if (Results.THROTTLED == result) {
			      retry = true;
			  } else if (Results.SERVER_ERROR == result) {
			      retry = true;
			  }
			  else {
			      // Some other error occurred, so stop calling the API.
			      retry = false;
			  }*/
			// SERVER_ERROR
			fmt.Printf("error requesting post: %v\n", err)
			break
		}

		fmt.Println(jsonplaceholder.Stringify(user))
		fmt.Println("=============================================================================")

		// stop processing if we are retry flag has
		if retries >= maxRetries {
			fmt.Printf("Reached Max Retries, exiting retries")
			break
		}

		// time.Sleep(1 * time.Second)
		backOffRequest(count, baseSleepFactor)
	}
}
