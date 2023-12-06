package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getAllWeatherByName(city string, api_key string) ([]byte, error) {

	request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, api_key)
	resp, err := http.Get(request)

	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %w\n", err)
	}

	defer resp.Body.Close()

	jsonData, _ := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func getPartialWeatherByName(city string, api_key string) ([]byte, error) {

	request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, api_key)
	resp, err := http.Get(request)

	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %w\n", err)
	}

	defer resp.Body.Close()

	type Weather struct {
		Description string `json:"description"`
	}

	type Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Humidity  int     `json:"humidity"`
	}

	type Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	}

	type Sys struct {
		Country string `json:"country"`
	}

	type Root struct {
		Weather []Weather `json:"weather"`
		Main    Main      `json:"main"`
		Wind    Wind      `json:"wind"`
		Sys     Sys       `json:"sys"`
		Name    string    `json:"name"`
	}

	var weatherInfo Root

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &weatherInfo) // get required subset of weather info
	if err != nil {
		return nil, err
	}

	jsonData, err = json.Marshal(weatherInfo)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
