package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"net/http"

	"github.com/converter/weather-report/api/openweather"
	"github.com/converter/weather-report/rest"
	"github.com/pkg/errors"
)

const (
	errNotEnoughArguments = iota + 1
	errSearchbyOptionMissing
	errSearchbyOptionUnknown
	errMarshalingWeather
)

type runOpts struct {
	searchBy    *string
	prettyPrint *bool
	term        *string
}

func main() {
	opts := runOpts{}
	opts.searchBy = flag.String("searchby", "", "openweather by city, zipcode or latlon")
	opts.prettyPrint = flag.Bool("pretty", false, "print nicely formatted output")
	flag.Parse()
	if *opts.searchBy == "" {
		usage("missing searchby option")
		os.Exit(errSearchbyOptionMissing)
	}
	switch *opts.searchBy {
	case "city":
	case "zipcode":
	case "latlon":
	default:
		usage("unknown searchby option: " + *opts.searchBy)
		os.Exit(errSearchbyOptionUnknown)
	}
	args := flag.Args()
	if len(args) < 1 {
		usage("not enough arguments")
		os.Exit(errNotEnoughArguments)
	}
	opts.term = &args[0]

	err := execCmd(opts)
	if err != nil {
		fmt.Printf("error %s", err.Error())
	}
}

func usage(msg string) {
	_, cmd := path.Split(os.Args[0])
	text := "%s\nusage: %s --searchby <type of term> <term>\n" +
		"example openweather terms:\n" +
		"city: New York, NY\n" +
		"zipcode: 74129\n" +
		"latlon: 38.505252,-90.430133\n"
	log.Printf(text, msg, cmd)
}

func getAPIKey() (string, error) {
	apikey := os.Getenv("WEATHER_API_KEY")
	if apikey != "" {
		return apikey, nil
	}

	b, err := ioutil.ReadFile(openweather.APIKeyFilename)

	if err != nil {
		log.Printf("error reading API key from %s: %s",
			openweather.APIKeyFilename, err.Error())
		return "", nil
	}
	return string(bytes.TrimSpace(b)), nil
}

func execCmd(opts runOpts) error {
	apikey, err := getAPIKey()
	if err != nil {
		log.Printf("")
	}
	var weather *openweather.OpenWeatherCurrent
	c := &rest.APIClient{HTTPClient: &http.Client{}}
	weather, err = c.GetWeather(apikey, *opts.searchBy, *opts.term)
	if err != nil {
		errors.Errorf("error fetching weather: %s", err.Error())
	}

	if *opts.prettyPrint {
		fmt.Println(weather.PrettyPrint())
		return nil
	}

	b, err := json.Marshal(weather)
	if err != nil {
		log.Printf("error marshaling weather data: %s", err.Error())
		os.Exit(errMarshalingWeather)
	}
	fmt.Println(string(b))

	return nil
}
