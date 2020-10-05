package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/converter/weather-report/api"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type APIClient struct {
	RequestURL string
}

func (c *APIClient) GetWeatherByCity(city string) (*api.OpenWeatherCurrent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.RequestURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &api.OpenWeatherCurrent{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read error: %s", err.Error())
	}

	weather := api.OpenWeatherCurrent{}

	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalf("could not unmarshal response JSON: %s", err.Error())
	}

	return &weather, nil
}
