package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

type MatchingResponse struct {
	Matches []string `json:"matches"`
}

func main() {
	http.HandleFunc("/api/match", matchHandler)
	log.Fatal(http.ListenAndServe(":8080", allowCORSHeaders(http.DefaultServeMux)))
}

func matchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		RegexPattern string `json:"regexPattern"`
		InputText    string `json:"inputText"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	re, err := regexp.Compile(requestData.RegexPattern)
	if err != nil {
		http.Error(w, "Invalid regex pattern", http.StatusBadRequest)
		return
	}

	matches := re.FindAllString(requestData.InputText, -1)

	response := MatchingResponse{
		Matches: matches,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// AllowCORSHeaders is a middleware that sets the necessary headers for Cross-Origin Resource Sharing (CORS)
func allowCORSHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers to allow cross-origin requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		handler.ServeHTTP(w, r)
	})
}
