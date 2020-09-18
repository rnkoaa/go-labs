package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		// server := NewServer(config)
		// sigChan, serveAndWait := server.Start()

		// // set the server to shutdown when SIGINT or SIGTERM is received
		// server.Shutdown(sigChan, serveAndWait)

		b, _ := json.Marshal(config)
		fmt.Printf("%s\n", b)
	},
}
