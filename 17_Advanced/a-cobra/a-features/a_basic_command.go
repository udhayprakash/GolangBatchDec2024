package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A simple CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from Cobra!")
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
