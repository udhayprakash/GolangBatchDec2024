package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func main() {
	var configFile string

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "An example CLI with config",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Config file:", configFile)
		},
	}

	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.app.yaml)")
	cobra.OnInitialize(initConfig)

	rootCmd.Execute()
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("$HOME")
		viper.SetConfigName(".app")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
}
