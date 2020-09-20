package main

import (
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		server := NewServer(config)
		sigChan, serveAndWait := server.Start()

		// set the server to shutdown when SIGINT or SIGTERM is received
		server.Shutdown(sigChan, serveAndWait)
	},
}
