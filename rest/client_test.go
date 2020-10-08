package rest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/converter/weather-report/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClient_GetWeather_ByCity(t *testing.T) {
	query := "q=tulsa,us-ok"
	mockHTTPClient := &mocks.MockHTTPClient{}
	matcher := func(req *http.Request) bool {
		if strings.Contains(req.URL.RequestURI(), query) {
			return true
		}
		return false
	}
	mockHTTPClient.On("Do", mock.MatchedBy(matcher)).Return(
		&http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(`{"Name":"Tulsa"}`)),
		},
		nil,
	)
	client := &APIClient{HTTPClient: mockHTTPClient}
	weather, err := client.GetWeather("DUMMYKEY", query)
	assert.NoError(t, err)
	assert.Equal(t, "Tulsa", weather.Name)
	mockHTTPClient.AssertExpectations(t)
}

func TestClient_GetWeather_ByZipCode(t *testing.T) {
	query := "zip=62221"
	mockHTTPClient := &mocks.MockHTTPClient{}
	matcher := func(req *http.Request) bool {
		if strings.Contains(req.URL.RequestURI(), query) {
			return true
		}
		return false
	}
	mockHTTPClient.On("Do", mock.MatchedBy(matcher)).Return(
		&http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(`{"Name":"Belleville"}`)),
		},
		nil,
	)
	client := &APIClient{HTTPClient: mockHTTPClient}
	weather, err := client.GetWeather("DUMMYKEY", query)
	assert.NoError(t, err)
	assert.Equal(t, "Belleville", weather.Name)
	mockHTTPClient.AssertExpectations(t)
}

func TestClient_GetWeather_ByLatLon(t *testing.T) {
	cityName := "Eastland Heights"
	query := "lat=36.185904&lon=-115.165393"
	mockHTTPClient := &mocks.MockHTTPClient{}
	matcher := func(req *http.Request) bool {
		if strings.Contains(req.URL.RequestURI(), query) {
			return true
		}
		return false
	}
	bod := `{"name":"Eastland Heights","coord":{"lon":-115.17,"lat":36.19}}`
	mockHTTPClient.On("Do", mock.MatchedBy(matcher)).Return(
		&http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(bod)),
		},
		nil,
	)
	client := &APIClient{HTTPClient: mockHTTPClient}
	weather, err := client.GetWeather("DUMMYKEY", query)
	assert.NoError(t, err)
	assert.Equal(t, cityName, weather.Name)
	assert.Equal(t, float32(36.19), weather.Coord.Lat)
	assert.Equal(t, float32(-115.17), weather.Coord.Lon)
	mockHTTPClient.AssertExpectations(t)
}

func TestClient_GetWeather_Pretty(t *testing.T) {
	expected := `Tulsa Weather:
Clear
Temp             73.0F
Feels like       66.7F
High             75.2F
Low              71.0F
Pressure      1020 hPa 
Humidity           41%
Wind          10.3 mph
Direction  170 degrees`
	jsonBody := `{
  "coord": {
    "lon": -95.9,
    "lat": 36.1
  },
  "weather": [
    {
      "id": 800,
      "main": "Clear",
      "description": "clear sky"
    }
  ],
  "base": "stations",
  "main": {
    "temp": 73,
    "feels_like": 66.74,
    "temp_min": 71.01,
    "temp_max": 75.2,
    "pressure": 1020,
    "humidity": 41
  },
  "visibility": 10000,
  "wind": {
    "speed": 10.29,
    "deg": 170
  },
  "clouds": {
    "all": 1
  },
  "dt": 1601926247,
  "sys": {
    "type": 1,
    "id": 5727,
    "country": "US",
    "sunrise": 1601900534,
    "sunset": 1601942471
  },
  "timezone": -18000,
  "id": 4553440,
  "name": "Tulsa",
  "cod": 200
}`
	term := "q=tulsa,us-ok"
	mockHTTPClient := &mocks.MockHTTPClient{}
	client := &APIClient{HTTPClient: mockHTTPClient}
	matcher := func(req *http.Request) bool {
		if strings.Contains(req.URL.RequestURI(), term) {
			return true
		}
		return false
	}
	mockHTTPClient.On("Do", mock.MatchedBy(matcher)).Return(
		&http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(jsonBody)),
		}, nil)
	weather, err := client.GetWeather("DUMMYKEY", term)
	assert.NoError(t, err)
	assert.Equal(t, expected, weather.PrettyPrint())
	mockHTTPClient.AssertExpectations(t)
}
