package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Name string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		data := PageData{Name: "Golang"}
		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
