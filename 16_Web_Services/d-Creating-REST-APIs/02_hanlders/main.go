package main

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	Name string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is ", h.Name)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World! dog")
}

func main() {

	// Method 1 - calling function handler
	http.HandleFunc("/dog", helloHandler)

	// Method 2 - calling method handler
	h := Handler{
		Name: "cat",
	}
	http.Handle("/cat", h)

	// Starting the server
	log.Println("Starting server at 8000 ...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
