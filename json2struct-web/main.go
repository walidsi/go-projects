package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/walidsi/json2struct"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file content
	htmlContent, err := os.ReadFile("jsonform.html")
	if err != nil {
		http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
		return
	}

	// Write the HTML content as the response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(htmlContent)
}

func handlerSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		// Get the value of the "textContent" field from the form
		jsonData := r.Form.Get("textContent")
		structString, err := json2struct.JSONToStruct("root", string(jsonData))

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(structString))
	}
	defer r.Body.Close()
}

func main() {
	// Register the handler function for the root endpoint ("/")
	http.HandleFunc("/", handler)
	http.HandleFunc("/submit", handlerSubmit)

	// Start the HTTP server on port 8080
	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
