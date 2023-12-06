package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// API Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world from Go http Server\n")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi\n")
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
