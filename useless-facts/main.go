package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type randomFact struct {
	Text string `json:"text"`
	ID   string `json:"id"`
}

func main() {
	fmt.Println("Hello World!")

	resp, err := http.Get("https://uselessfacts.jsph.pl/api/v2/facts/random")

	if err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
	}

	defer resp.Body.Close()

	var fact randomFact
	json.NewDecoder(resp.Body).Decode(&fact)
	fmt.Println(fact)
}
