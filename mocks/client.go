package mocks

import (
	"log"

	"github.com/converter/weather-report/api"
)

type MockAPIClient struct {
	RequestURL string
}

func (c *MockAPIClient) GetWeatherByCity(city string) (*api.OpenWeatherCurrent, error) {
	log.Printf("mock http client called.")
	return &api.OpenWeatherCurrent{}, nil
}
