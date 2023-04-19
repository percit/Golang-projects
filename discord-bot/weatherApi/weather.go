package weatherApi

import (
	"net/http"
	"encoding/json"
	"time"
	"fmt"
)

const openWeatherMapURL = "http://api.openweathermap.org/data/2.5/weather"

type OpenWeatherMapResponse struct {
	Name      string `json:"name"`
	Main      Main   `json:"main"`
	Weather   []Weather `json:"weather"`
	Timestamp int64     `json:"dt"`
}

type Main struct {
	Temperature float64 `json:"temp"`
}

type Weather struct {
	Description string `json:"description"`
}


func GetWeather(city string, apiKey string) (string, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", openWeatherMapURL, city, apiKey)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	var openWeatherMapResponse OpenWeatherMapResponse

	if err := json.NewDecoder(resp.Body).Decode(&openWeatherMapResponse); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	weather := openWeatherMapResponse.Weather[0].Description
	temperature := fmt.Sprintf("%.1f", openWeatherMapResponse.Main.Temperature)
	cityName := openWeatherMapResponse.Name

	return fmt.Sprintf("Current weather in %s: %s, Temperature: %s℃", cityName, weather, temperature), nil
}