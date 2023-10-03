package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"openweather-api-http-server-example/alert"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var OpenWeatherAPIKey = os.Getenv("OWM_API_KEY")

type Coordinates struct {
	Lat string `form:"lat"`
	Lon string `form:"lon"`
}

func main() {
	router := gin.Default()
	router.GET("/weather", getCurrentWeather)
	router.Run(":8080")
}

func getCurrentWeather(c *gin.Context) {
	var coords Coordinates
	if c.BindQuery(&coords) == nil {
		log.Printf("lat: %s lon: %s", coords.Lat, coords.Lon)
	}
	weatherData, err := fetchCurrentWeatherData(coords.Lat, coords.Lon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	condition := weatherData.Weather[0].Description
	temperature := weatherData.Main.Temp

	var temperatureFeel string
	if temperature < 55 {
		temperatureFeel = "cold"
	} else if temperature >= 55 && temperature <= 80 {
		temperatureFeel = "moderate"
	} else {
		temperatureFeel = "hot"
	}

	// NOTE:The *free* OpenWeather API endpoint for current weather does not include alerts
	// TODO: source from NWS free open api if time allows
	activeAlertsData, err := fetchActiveAlertsData(coords.Lat, coords.Lon)
	validActiveAlerts := alert.ProcessNWSActiveAlertResponse(activeAlertsData)

	c.JSON(http.StatusOK, gin.H{
		"condition":        condition,
		"temperature_feel": temperatureFeel,
		"activeAlerts":     validActiveAlerts,
	})
}

func fetchActiveAlertsData(lat string, lon string) (*alert.NWSActiveAlertsResponse, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("point", fmt.Sprintf("%s,%s", lat, lon)).
		SetDebug(true).
		Get("https://api.weather.gov/alerts/active")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch Active Alerts data: %s", resp.Status())
	}

	var activeAlertsData alert.NWSActiveAlertsResponse
	if err := json.Unmarshal(resp.Body(), &activeAlertsData); err != nil {
		log.Print("Unable to unmarshal Active Alerts response body: ", resp.Body())
		return nil, err
	}

	return &activeAlertsData, nil
}

func fetchCurrentWeatherData(lat string, lon string) (*CurrentWeatherResponse, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("lat", lat).
		SetQueryParam("lon", lon).
		SetQueryParam("appid", OpenWeatherAPIKey).
		SetQueryParam("units", "imperial").
		SetDebug(true).
		Get("https://api.openweathermap.org/data/2.5/weather")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data: %s", resp.Status())
	}

	var weatherData CurrentWeatherResponse
	if err := json.Unmarshal(resp.Body(), &weatherData); err != nil {
		log.Print("Unable to unmarshal response body: ", resp.Body())
		return nil, err
	}

	return &weatherData, nil
}
