package main

func main() {
	Execute()

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// cr := cron.New()
	// _, err := cr.AddFunc("0/1 * * * ?", func() {
	// 	currentTime := time.Now().Format("2006-01-02T15:04:05-0700")
	// 	fmt.Printf("[%s] Starting cron job\n", currentTime)
	// })
	// if err != nil {
	// 	log.Fatalf("error adding function to cron")
	// }

	// select {
	// case <-sigChan:
	// 	log.Printf("Shutdown signal received... closing server")
	// 	cr.Stop()
	// default:
	// 	currentTime := time.Now().Format("2006-01-02T15:04:05-0700")
	// 	fmt.Printf("[%s] Starting cron job\n", currentTime)
	// 	cr.Run()
	// }

}
