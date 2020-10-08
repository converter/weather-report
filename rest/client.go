package rest

import (
	"context"
	"encoding/json"
	"fmt"
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

func (c *APIClient) GetWeather(apikey, query string) (*openweather.OpenWeatherCurrent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	requestURL := fmt.Sprintf(openweather.URLPattern, query, apikey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
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
