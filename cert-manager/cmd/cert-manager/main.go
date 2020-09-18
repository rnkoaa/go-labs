package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	// "time"
	// "github.com/go-co-op/gocron"
)

func main() {
	// Execute()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// defines a new scheduler that schedules and runs jobs
	// s1 := gocron.NewScheduler(time.UTC)
	// s1.Every(3).Seconds().Do(task)

	select {
	case <-sigChan:
		log.Printf("Shutdown signal received... closing server")
		// s1.Stop()
	default:
		// <-s1.StartAsync()
	}

}
