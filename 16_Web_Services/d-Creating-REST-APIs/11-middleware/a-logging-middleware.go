package main

import (
	"fmt"
	"net/http"
)

/*
Middleware is an entity that intercepts the server's request/response life cycle.
  - Process the request before running business logic (authentication)
  - Modify the request to the next handler function (attaching payload)
  - Modify the response for the client
  - Logging.... and much more
*/

// Middleware to log requests
func loggingMiddleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before handler")
		w.Write([]byte("Hijacking Request "))
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Running after handler")
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "<h1>Hello!</h1>")
	fmt.Println("Executing the handler")
	w.Write([]byte("OK"))
}

func main() {
	http.Handle("/hello", loggingMiddleware(http.HandlerFunc(helloHandler)))

	fmt.Println("Listening at localhost:3000")
	http.ListenAndServe(":3000", nil)
}
