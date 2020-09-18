package main

import (
	"fmt"
)

// CertManager - struct to manage certs
type CertManager struct {
	config *Config
	// cronScheduler *cron.Cron
}

func NewCertManager(config *Config) *CertManager {
	return &CertManager{
		config: config,
	}
}

func (c CertManager) Start() error {
	// c.cronScheduler.AddFunc(c.config.cronSchedule, func() {
	// 	fmt.Println("Chron executed")
	// })

	// if err != nil {
	// 	return err
	// }
	// c.cronScheduler.Start()

	// defines a new scheduler that schedules and runs jobs
	// s1 := gocron.NewScheduler(time.UTC)

	// s1.Every(3).Seconds().Do(task)

	// scheduler starts running jobs and current thread continues to execute
	// s1.StartBlocking()
	return nil
}

func task() {
	fmt.Println("I am running task.")
}

func (c CertManager) Stop() {
	// c.cronScheduler.Stop()
}
