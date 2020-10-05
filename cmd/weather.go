package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/converter/weather-report/rest"
)

type runOpts struct {
	searchBy *string
	term     *string
}

func main() {
	opts := runOpts{}
	opts.searchBy = flag.String("searchby", "", "search by city, zip or latlong")
	flag.Parse()
	switch *opts.searchBy {
	case "city":
	case "zip":
	case "latlong":
	default:
		usage("unknown searchby option: " + *opts.searchBy)
		os.Exit(1)
	}
	args := flag.Args()
	if len(args) < 1 {
		usage("not enough arguments")
		os.Exit(1)
	}
	opts.term = &args[0]

	execCmd(opts)
}

func usage(msg string) {
	_, cmd := path.Split(os.Args[0])
	text := "%s\nusage: %s --searchby <type of term> <term>\n" +
		"example search terms:\n" +
		"city: New York, NY\n" +
		"zip: 74129\n" +
		"latlong: 38.505252,-90.430133\n"
	log.Printf(text, msg, cmd)
}

func execCmd(opts runOpts) {
	filename := "./.weather-api-key"
	b, err := ioutil.ReadFile(filename)
	apikey := string(bytes.TrimSpace(b))
	if err != nil {
		log.Fatalf("error reading API key from %s: %s", filename, err.Error())
	}
	cityst := "Fenton,US-MO"
	u := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial&APPID=%s",
		cityst, apikey)
	c := &rest.APIClient{RequestURL: u}
	weather, err := c.GetWeatherByCity("Fenton,US-MO")
	if err != nil {
		log.Fatalf("")
	}
	log.Printf("weather = %#v", weather)
}
