package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	http.HandleFunc("/hello/", hello)
	http.ListenAndServe("localhost:8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("req.URL.Path", req.URL.Path)

	var helloExp = regexp.MustCompile(`/hello/(?P<name>[a-zA-Z]+)/(?P<age>\d+)`)
	match := helloExp.FindStringSubmatch(req.URL.Path)
	if len(match) > 0 {
		result := make(map[string]string)
		for i, name := range helloExp.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		if _, err := strconv.Atoi(result["age"]); err == nil {
			fmt.Fprintf(w, "Hello, %v year old named %s!", result["age"], result["name"])
		} else {
			fmt.Fprintf(w, "Sorry, not accepted age!")
		}
	} else {
		fmt.Fprintf(w, "Wrong url\n")
	}
}

/*
/hello/(?P<name>[a-zA-Z]+)/(?P<age>\d+)


(?P<name>[a-zA-Z]+)
		[a-zA-Z]+     lower, upper case alphanets

(?P<age>\d+)
		\d+			only digits, as number


http://localhost:8090/hello/RobPike/56
http://localhost:8090/hello/JOHN/56
http://localhost:8090/hello/john/56

http://localhost:8090/hello/john34/56     wrong url
http://localhost:8090/hello/john/three	  wrong url
*/
