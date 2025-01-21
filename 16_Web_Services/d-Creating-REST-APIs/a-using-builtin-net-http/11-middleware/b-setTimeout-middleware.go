package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)


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

func main() {
	http.Handle("/hello", setTimeCookie(http.HandlerFunc(helloHandler)))

	fmt.Println("Listening at localhost:3000")
	http.ListenAndServe(":3000", nil)
}
