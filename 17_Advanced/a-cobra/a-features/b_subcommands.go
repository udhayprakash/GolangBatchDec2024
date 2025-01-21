package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	var helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints Hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello!")
		},
	}

	rootCmd.AddCommand(helloCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// - go run b_subcommands.go hello
