package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	secretFile string
	config     *Config
	rootCmd    = &cobra.Command{
		Use: "cert-manager",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Usage()
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file location")
	rootCmd.PersistentFlags().StringVarP(&secretFile, "secret", "s", "", "secret file location")
	rootCmd.AddCommand(serveCmd)
}

func initConfig() {
	config = &Config{}
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yml")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("unable to read config using path: ", err)
	}

	if secretFile != "" {
		viper.SetConfigFile(secretFile)
		viper.SetConfigType("yml")
		if err := viper.MergeInConfig(); err != nil {
			fmt.Println("unable to read secret file: ", err)
		}
	}

	viper.AutomaticEnv()

	viper.SetDefault("cron.schedule", "@daily")
	viper.SetDefault("address", "localhost")
	viper.SetDefault("port", "8080")

	err := viper.Unmarshal(config)
	if err != nil {
		log.Fatalln("error unmarshalling viper into config ", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("err while starting application, %v\n", err)
		os.Exit(1)
	}
}
