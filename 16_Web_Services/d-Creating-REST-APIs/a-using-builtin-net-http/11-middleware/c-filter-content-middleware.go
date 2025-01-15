package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)


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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "<h1>Hello!</h1>")
	fmt.Println("Executing the handler")
	w.Write([]byte("OK"))
}

func main() {
	http.Handle("/hello", filterContentType(http.HandlerFunc(helloHandler)))

	fmt.Println("Listening at localhost:3000")
	http.ListenAndServe(":3000", nil)
}
