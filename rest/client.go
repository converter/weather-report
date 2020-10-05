package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"strings"

	"github.com/converter/weather-report/api/openweather"
)

type HTTPer interface {
	Do(req *http.Request) (*http.Response, error)
}

type APIClient struct {
	HTTPClient HTTPer
}

func (c *APIClient) GetWeatherByCity(apikey, query string) (*openweather.OpenWeatherCurrent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query = "q=" + query
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

func (c *APIClient) GetWeatherByZipCode(apikey, query string) (*openweather.OpenWeatherCurrent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query = "zip=" + query
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

func (c *APIClient) GetWeatherByLatLon(apikey, query string) (*openweather.OpenWeatherCurrent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query = strings.TrimSpace(query)
	tokens := strings.Split(query, ",")
	if len(tokens) != 2 {
		log.Fatalf("invalid lat lon: %s", query)
	}

	query = fmt.Sprintf("lat=%s&lon=%s", tokens[0], tokens[1])
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
