package main

import (
	"fmt"
	"log"
	"net/http"
)

func requestHeaders(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%25v: %v\n", name, h)
		}
	}
}

func setHeaders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Custom-Header", "CustomValue")
	w.Write([]byte(`{"message": "Headers set"}`))
}

func main() {
	http.HandleFunc("/getHeader", requestHeaders)
	http.HandleFunc("/setHeader", setHeaders)

	log.Fatal(http.ListenAndServe(":9001", nil))
}


/*

/getHeader

          Accept-Encoding: gzip, deflate, br
               Connection: keep-alive
               User-Agent: PostmanRuntime/7.43.0
                   Accept: *
				   Postman-Token: 98dfb010-4fd1-4b56-8dc0-1e33a68e713b


setHeader
 {
    "message": "Headers set"
}

*/