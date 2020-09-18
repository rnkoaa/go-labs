package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	config  *Config
	rootCmd = &cobra.Command{
		Use: "cert-manager",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Usage()
		},
	}
)

func init() {
	config = &Config{
		cronSchedule: "*/1 * * * *",
		Address:      "localhost",
		Port:         "8080",
	}
	rootCmd.AddCommand(serveCmd)
}

func Execute() {
	// server := NewServer(config)
	// server.Start()

	// server.Shutdown()
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("err while starting application, %v\n", err)
		os.Exit(1)
	}
}
