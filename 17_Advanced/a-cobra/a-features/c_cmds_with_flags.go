package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var name string

	var greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "Greet someone",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	greetCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*

- go run c_cmds_with_flags.go greet -n Golang
Hello, Golang!

- go run c_cmds_with_flags.go greet --name Golang
Hello, Golang!

- go run c_cmds_with_flags.go greet
Hello, World!


*/
