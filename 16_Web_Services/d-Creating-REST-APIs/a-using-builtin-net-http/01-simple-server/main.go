package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHanlder(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "default endpoint")
}

func helloHanlder(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World")
}

func main() {
	http.HandleFunc("/", defaultHanlder)
	http.HandleFunc("/hello", helloHanlder)

	log.Println("Starting server at 8000 ...")

	// Starting server - default is localhost
	// log.Fatal(http.ListenAndServe(":8000", nil))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
