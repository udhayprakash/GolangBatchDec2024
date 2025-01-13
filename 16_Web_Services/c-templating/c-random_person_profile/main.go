package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"text/template"
)

type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Quote      string `json:"quote"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// step 1 - reading all data, from file
		data, err := ioutil.ReadFile("data/persons.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parsing and placing in slice
		var persons []Person
		err = json.Unmarshal(data, &persons)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Step 3: Choosing a random person
		randomPerson := persons[rand.Intn(len(persons))]

		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/profile.html"))
		if err := tmpl.ExecuteTemplate(w, "profile.html", randomPerson); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

/*
Assignment 
==========
	create a random person details by taking 
		surnames separately in a text file
		male names separately in a text file
		female names separately in a text file

	Then, combine them and create random profile info. 
*/