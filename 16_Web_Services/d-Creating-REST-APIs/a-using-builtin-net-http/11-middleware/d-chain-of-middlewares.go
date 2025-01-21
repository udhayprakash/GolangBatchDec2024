package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Middleware to log requests
func loggingMiddleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before handler")
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Running after handler")
	})
}

// Middleware to filter content type
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("405 - Header Content-Type incorrect"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Middleware to set a cookie with the server time
func setTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:  "Server-Time-UTC",
			Value: strconv.Itoa(int(time.Now().Unix())),
		}
		http.SetCookie(w, &cookie)
		handler.ServeHTTP(w, r)
	})
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "<h1>Hello!</h1>")
	fmt.Println("Executing the handler")
	w.Write([]byte("OK"))
}

// Chain middlewares together
type Middleware func(http.Handler) http.Handler

func chainMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}


func main() {
	// // Case 1: Using individual middleware
	// http.Handle("/hello", loggingMiddleware(http.HandlerFunc(helloHandler)))

	// Case 2: Using a combination of middlewares
	combinedMiddlewares := chainMiddlewares(
		http.HandlerFunc(helloHandler),
		loggingMiddleware,
		filterContentType,
		setTimeCookie,
	)
	http.Handle("/helloWithMiddlewares", combinedMiddlewares)

	fmt.Println("Listening at localhost:3000")
	http.ListenAndServe(":3000", nil)
}