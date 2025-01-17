package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "An example CLI",
		Long:  "This is a long description for the example CLI application.",
	}

	var cmd = &cobra.Command{
		Use:   "run",
		Short: "Run the application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running the app...")
		},
	}

	rootCmd.AddCommand(cmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
