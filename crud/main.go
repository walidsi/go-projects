package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/walidsi/go-projects/crud/internal/application"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables")
	}
}

func main() {
	log.Fatal(application.Run())
}
