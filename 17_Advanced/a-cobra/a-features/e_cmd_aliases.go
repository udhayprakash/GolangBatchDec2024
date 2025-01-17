package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	var byeCmd = &cobra.Command{
		Use:     "bye",
		Aliases: []string{"goodbye"},
		Short:   "Say goodbye",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Goodbye!")
		},
	}

	rootCmd.AddCommand(byeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*

- go run e_cmd_aliases.go bye
Goodbye!

- go run e_cmd_aliases.go goodbye
Goodbye!


*/
