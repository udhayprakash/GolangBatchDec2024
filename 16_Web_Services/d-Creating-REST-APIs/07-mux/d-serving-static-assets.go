package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexEndpoint)

	staticHandler := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))

	// Starting the server
	log.Println("Starting server at 3000 ...")
	log.Fatal(http.ListenAndServe(":3000", mux)) // passing userdefined mux
}

func IndexEndpoint(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Index endpoint")
}
