package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"tempC"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func main() {
	q := "Bulk"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=7b8feb261f1f4c80919112629231809&q=" + q)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode != 200 {
		panic("Weather API no available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return
	}
	location, current := weather.Location, weather.Current
	message, _ := fmt.Printf(
		"%s-%s: %.0fC, %s ",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)
	if current.TempC < 0 {
		fmt.Print(message)
	} else {
		color.Red(strconv.Itoa(message))
	}
}
