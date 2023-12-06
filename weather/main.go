package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func getInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	return input
}

func getCity() string {
	city := getInput("Enter city name (q to quit): ")

	return city
}

func getChoice() string {
	choice := getInput("Do you want (f)ull or (p)artial weather info?: ")

	switch choice {
	case "f":
	case "p":
		break
	default:
		fmt.Println("Unknown choice, please try again.")
		return getChoice()
	}
	return choice
}

func formatAndPrintJson(jsonData []byte) error {
	var prettyJSON bytes.Buffer

	err := json.Indent(&prettyJSON, jsonData, "", "  ")

	if err != nil {
		return err
	}

	fmt.Println(prettyJSON.String())

	return nil
}

func run(apiKey string) error {
	city := getCity()

	if city == "q" {
		return nil
	}

	choice := getChoice()

	var jsonInfo []byte

	switch choice {
	case "f":
		jsonInfo, _ = getAllWeatherByName(city, apiKey)
		formatAndPrintJson(jsonInfo)
	case "p":
		jsonInfo, _ = getPartialWeatherByName(city, apiKey)
		formatAndPrintJson(jsonInfo)
	}

	run(apiKey)

	return nil
}

func main() {
	fmt.Println("Welcome to Weather CLI!")

	var apiKey = goDotEnvVariable("OWM_API_KEY")

	run(apiKey)
}
