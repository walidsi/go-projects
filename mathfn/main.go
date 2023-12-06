package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodGet {
		r.ParseForm()
		a, err1 := strconv.Atoi(r.FormValue("a"))
		b, err2 := strconv.Atoi(r.FormValue("b"))
		if err1 != nil || err2 != nil {
			http.Error(w, "Error converting string to integer", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Received a: %v and b: %v, sum = %v\n", a, b, a+b)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func subtract(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	a, err1 := strconv.Atoi(r.FormValue("a"))
	b, err2 := strconv.Atoi(r.FormValue("b"))
	if err1 != nil || err2 != nil {
		http.Error(w, "Error converting string to integer", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received a: %v and b: %v, difference = %v\n", a, b, a-b)
}

func multiply(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	a, err1 := strconv.Atoi(r.FormValue("a"))
	b, err2 := strconv.Atoi(r.FormValue("b"))
	if err1 != nil || err2 != nil {
		http.Error(w, "Error converting string to integer", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received a: %v and b: %v, multiplication = %v\n", a, b, a*b)
}

func divide(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	a, err1 := strconv.ParseFloat(r.FormValue("a"), 64)
	b, err2 := strconv.ParseFloat(r.FormValue("b"), 64)

	if err1 != nil || err2 != nil {
		http.Error(w, "Error converting string to integer", http.StatusBadRequest)
		return
	}
	if b == 0 {
		http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received a: %v and b: %v, division = %0.4f\n", a, b, a/b)
}

func main() {
	// API Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the math REST API\n")
	})

	http.HandleFunc("/add/", add)
	http.HandleFunc("/subtract/", subtract)
	http.HandleFunc("/multiply/", multiply)
	http.HandleFunc("/divide/", divide)

	port := ":5000"
	fmt.Println("Server is running on http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
