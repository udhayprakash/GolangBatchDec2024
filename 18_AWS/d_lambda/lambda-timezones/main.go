package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type TimeZone struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

type Response struct {
	TimeZones []TimeZone `json:"timezones"`
}

func Handler(ctx context.Context) (Response, error) {
	locations := []string{
		"UTC",
		"America/New_York",
		"Europe/London",
		"Asia/Tokyo",
		"Australia/Sydney",
		"Pacific/Honolulu",
		"Asia/Kolkata",
		"Europe/Paris",
		"America/Los_Angeles",
		"Asia/Shanghai",
	}

	var timeZones []TimeZone
	for _, loc := range locations {
		location, err := time.LoadLocation(loc)
		if err != nil {
			return Response{}, err
		}
		currentTime := time.Now().In(location)
		timeZones = append(timeZones, TimeZone{
			Name: loc,
			Time: currentTime.Format(time.RFC3339),
		})
	}

	return Response{TimeZones: timeZones}, nil
}

func main() {
	lambda.Start(Handler)
}
