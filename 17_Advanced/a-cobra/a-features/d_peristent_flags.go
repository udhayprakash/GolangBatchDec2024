package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var verbosity int

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.PersistentFlags().IntVarP(&verbosity, "verbosity", "v", 1, "Verbosity level")

	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "Show info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Verbosity level: %d\n", verbosity)
		},
	}

	rootCmd.AddCommand(infoCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*

- go run d_peristent_flags.go info
Verbosity level: 1

- go run d_peristent_flags.go info --verbosity 2
Verbosity level: 2


*/
