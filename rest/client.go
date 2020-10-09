package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/converter/weather-report/api/openweather"
)

type HTTPer interface {
	Do(req *http.Request) (*http.Response, error)
}

type APIClient struct {
	HTTPClient HTTPer
}

// GetWeather makes an HTTP request to the OpenWeather API and returns an OpenWeather struct/object or an error.
func (c *APIClient) GetWeather(apiKey, searchBy, term string) (*openweather.OpenWeatherCurrent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	uri, err := openweather.ComposeRequestURI(apiKey, searchBy, term)
	if err != nil {
		return &openweather.OpenWeatherCurrent{}, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return &openweather.OpenWeatherCurrent{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read error: %s", err.Error())
	}

	weather := openweather.OpenWeatherCurrent{}
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalf("could not unmarshal response JSON: %s", err.Error())
	}

	return &weather, nil
}
