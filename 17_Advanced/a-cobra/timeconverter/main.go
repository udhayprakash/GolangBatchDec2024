package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "time-converter [time] [timezone]",
		Short: "Convert local time to UTC",
		Long:  `Convert a given local time in any timezone to UTC, considering daylight savings.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			input := strings.Join(args, " ")
			utcTime, err := convertToUTC(input)
			if err != nil {
				log.Fatalf("Error converting time: %s", err)
			}
			fmt.Println(utcTime)
		},
	}

}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func convertToUTC(input string) (string, error) {
	// Define a map for timezone abbreviations considering daylight savings
	timeZoneMap := map[string]string{
		"EST": "America/New_York",
		"EDT": "America/New_York",
		"CST": "America/Chicago",
		"CDT": "America/Chicago",
		"MST": "America/Denver",
		"MDT": "America/Denver",
		"PST": "America/Los_Angeles",
		"PDT": "America/Los_Angeles",
	}

	// Parse the input time
	layout := "3:04 PM"
	t, err := time.Parse(layout, input)
	if err != nil {
		return "", err
	}

	// Extract timezone from input
	var tz string
	for abbr := range timeZoneMap {
		if strings.Contains(input, abbr) {
			tz = timeZoneMap[abbr]
			break
		}
	}

	if tz == "" {
		return "", fmt.Errorf("timezone not recognized")
	}

	// Load location
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return "", err
	}

	// Convert to UTC
	utcTime := t.In(loc).UTC()

	// Format and return the UTC time
	return utcTime.Format("03:04 PM UTC"), nil
}
