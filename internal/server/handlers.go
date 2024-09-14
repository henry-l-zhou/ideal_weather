package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) getWeather(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	weatherData, err := s.weatherClient.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := struct {
		City        string  `json:"city"`
		Temperature float64 `json:"temperature"`
		Description string  `json:"description"`
	}{
		City:        city,
		Temperature: weatherData.Main.Temp - 273.15,
		Description: "",
	}

	if len(weatherData.Weather) > 0 {
		result.Description = weatherData.Weather[0].Description
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
