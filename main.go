package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ManishBadgotra/weather-cli/models"
	"github.com/joho/godotenv"
)

var (
	ApiKey  string
	BaseURL string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var statusErr error

	fmt.Println("Weather CLI")

	ApiKey = os.Getenv("apikey")
	if ApiKey == "" {
		log.Fatal("Api key did not load\n")
	}

	BaseURL = os.Getenv("baseURL")
	if ApiKey == "" {
		log.Fatal("Api key did not load\n")
	}

	for t := range time.Tick(time.Second) {

		timeString := t.String()
		url := fmt.Sprintf("%v?lat=28.611665&lon=76.978678&units=metric&date=%v&appid=%v", BaseURL, timeString[:11], ApiKey)

		httpResp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: ", err, " while calling an weather api")
			return
		}

		bytesResponse, err := io.ReadAll(httpResp.Body)
		if err != nil {
			fmt.Println("Error in reading data from a Response of a api")
			return
		}

		if httpResp.StatusCode != int(200) {
			errString := fmt.Sprintf("error %v while callback to endpoint", httpResp.StatusCode)
			statusErr = errors.New(errString)

		} else {
			var data models.WeatherAPI
			err = json.Unmarshal(bytesResponse, &data)
			if err != nil {
				fmt.Println("unable to Unpack data into Weather Structure.")
				return
			}

			fmt.Printf("Latitude: %v\n", data.Lat)
			fmt.Printf("Longitude: %v\n", data.Lon)
			fmt.Printf("Time Zone: %v\n", data.Timezone)
			fmt.Printf("Time Zone offset: %v\n", data.TimezoneOffset)
			fmt.Printf("Cloud: %v%%\n", data.Current.Clouds)
			fmt.Printf("Humidity: %v%%\n", data.Current.Humidity)
			fmt.Printf("Pressure: %v%%\n", data.Current.Pressure)
			degSymbol := '\u00B0'
			fmt.Printf("Temprature: %v%cC\n", data.Current.Temp, degSymbol)

			fmt.Println()
		}

		if statusErr != nil {
			fmt.Println(statusErr)
		}
	}
}
