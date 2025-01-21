package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	var failCmd = &cobra.Command{
		Use:   "fail",
		Short: "Fails intentionally",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This command will fail!")
			os.Exit(1) // Simulate failure
		},
	}

	rootCmd.AddCommand(failCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
