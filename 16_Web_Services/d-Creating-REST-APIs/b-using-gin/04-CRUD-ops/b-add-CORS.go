package main

import "net/http"

/*


server.go
client.go (postman, curl command, another go progam, another langauage program)
   same machine
 CORS - CORS (Cross-Origin Resource Sharing)

major  CORS Headers:
--------------------
1) Access-Control-Allow-Origin:
	Specifies which origins are allowed to access the resource.

	Example: Access-Control-Allow-Origin: https://example.com
			 Access-Control-Allow-Origin: * (to allow all origins).

2) Access-Control-Allow-Methods:

	Specifies which HTTP methods are allowed (e.g., GET, POST, PUT, DELETE).
	Example: Access-Control-Allow-Methods: GET, POST, OPTIONS

3) Access-Control-Allow-Headers:

	Specifies which headers can be used in the actual request.
	Example: Access-Control-Allow-Headers: Content-Type, Authorization.

4) Access-Control-Allow-Credentials:

	Indicates whether the browser should include credentials (e.g., cookies, tokens) in the request.
	Example: Access-Control-Allow-Credentials: true.
*/

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow specific methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		// Allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create a simple API handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, CORS!"))
	})

	// Wrap the handler with CORS middleware
	http.Handle("/", enableCORS(handler))

	// Start the server
	http.ListenAndServe(":8080", nil)
}
