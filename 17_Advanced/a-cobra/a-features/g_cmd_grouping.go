package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	var mathCmd = &cobra.Command{Use: "math"}
	var addCmd = &cobra.Command{
		Use:   "add [a] [b]",
		Short: "Add two numbers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Result: %s + %s\n", args[0], args[1])
		},
	}

	mathCmd.AddCommand(addCmd)
	rootCmd.AddCommand(mathCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
- go run g_cmd_grouping.go math  add 5 10
Result: 5 + 10


*/
