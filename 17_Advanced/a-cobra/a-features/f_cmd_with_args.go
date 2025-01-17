package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo a message",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(echoCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*


- go run f_cmd_with_args.go echo "Hello Wordl""
Hello Wordl"

- go run f_cmd_with_args.go echo "Hello"
Hello

- go run f_cmd_with_args.go echo Hello
Hello


*/
