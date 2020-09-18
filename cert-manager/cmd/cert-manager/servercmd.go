package main

import (
	"context"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		// server := NewServer(config)
		// sigChan, serveAndWait := server.Start()
		// server.Shutdown(sigChan, serveAndWait)

		// test chron work
		// config := &Config{
		// }
		certManager := NewCertManager(config)

		// signalChan := make(chan os.Signal, 1)
		// signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)

		// ctx, _ := context.WithCancel(context.Background())
		// defer func() {
		// 	signal.Stop(signalChan)
		// 	// cancel()
		// }()

		// go func() {
		// 	for {
		// 		select {
		// 		case s := <-signalChan:
		// 			switch s {
		// 			case os.Interrupt:
		// 				// cancel()
		// 				os.Exit(1)
		// 			}
		// 			// case <-ctx.Done():
		// 			// 	log.Printf("Done.")
		// 			// 	os.Exit(1)
		// 		}
		// 	}
		// }()
		// run(ctx, certManager)
		certManager.Start()
	},
}

func run(ctx context.Context, certManager *CertManager) {
	select {
	case <-ctx.Done():
		// certManager.Stop()
		// return
	default:
		certManager.Start()
	}
}
