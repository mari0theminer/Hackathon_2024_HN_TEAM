package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const APIKeyHeader = "X-API-Key"
const ExpectedAPIKey = "your_expected_api_key_here"

type InputData struct {
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
	Value  string  `json:"value"`
}

type Label struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get(APIKeyHeader)
		if apiKey != ExpectedAPIKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func InputHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var inputData InputData
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Here you would handle the inputData, e.g., save it to a database or process it

	response := ResponseMessage{Message: "Data successfully received."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/input", APIKeyMiddleware(http.HandlerFunc(InputHandler)))

	fmt.Println("Server is running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
