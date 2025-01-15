package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/bad-request", badRequestHandler)
	http.HandleFunc("/unauthorized", unauthorizedHandler)
	http.HandleFunc("/forbidden", forbiddenHandler)
	http.HandleFunc("/not-found", notFoundHandler)
	http.HandleFunc("/internal-server-error", internalServerErrorHandler)

	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)
}

// 400 Bad Request
func badRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "400 - Bad Request: Invalid input provided.",
	})
}

// 401 Unauthorized
func unauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "401 - Unauthorized: Authentication required.",
	})
}

// 403 Forbidden
func forbiddenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "403 - Forbidden: You do not have permission to access this resource.",
	})
}

// 404 Not Found
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "404 - Not Found: The requested resource could not be found.",
	})
}

// 500 Internal Server Error
func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response{
		Status:  "error",
		Message: "500 - Internal Server Error: Something went wrong on the server.",
	})
}
