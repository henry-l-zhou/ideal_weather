package goweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	APIKey  string
	BaseURL string
}

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "http://api.openweathermap.org/data/2.5/weather",
	}
}

func (c *Client) GetWeather(city string) (*WeatherData, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s", c.BaseURL, city, c.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error from weather API: %d", resp.StatusCode)
	}

	var data WeatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error parsing weather data: %w", err)
	}

	return &data, nil
}
