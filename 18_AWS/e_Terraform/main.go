package main

import (
	"context"
	"fmt"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (string, error) {
	if event == nil {
		return "Hello, World!", nil
	}
	return fmt.Sprintf("Hello, %s!", event.Name), nil
}


func main() {}
