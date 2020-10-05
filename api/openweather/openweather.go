package openweather

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"
)

const (
	APIKeyFilename = "./.weather-api-key"
	URLPattern     = "http://api.openweathermap.org/data/2.5/weather?%s&units=imperial&APPID=%s"
)

type OpenWeatherCurrent struct {
	Coord      `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       `json:"main"`
	Visibility int `json:"visibility"`
	Wind       `json:"wind"`
	Clouds     `json:"clouds"`
	DT         int32 `json:"dt"`
	Sys        `json:"sys"`
	Timezone   int    `json:"timezone"`
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	Cod        int    `json:"cod"`
}
type Coord struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}
type Main struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}
type Wind struct {
	Speed float32 `json:"speed"`
	Deg   int     `json:"deg"`
}
type Clouds struct {
	All int `json:"all"`
}
type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int32  `json:"sunrise"`
	Sunset  int32  `json:"sunset"`
}

type units struct {
	Temp      string
	WindSpeed string
}

type tmplData struct {
	*OpenWeatherCurrent
	Units units
}

// PrettyPrint uses a template to print nicely formatted weather output.
func (w *OpenWeatherCurrent) PrettyPrint() string {
	tmpl := `{{.Name}} Weather:
{{ range .Weather }}{{.Main}}{{end}}
Temp       {{ Temperature .Main.Temp .Units.Temp }}
Feels like {{ Temperature .Main.FeelsLike .Units.Temp }}
High       {{ Temperature .Main.TempMax .Units.Temp }}
Low        {{ Temperature .Main.TempMin .Units.Temp }}
Pressure   {{ Int .Main.Pressure " hPa" }} 
Humidity   {{ Int .Main.Humidity "%" }}
Wind       {{ WindSpeed .Wind.Speed .Units.WindSpeed }}
Direction  {{ Int .Wind.Deg " degrees" }}`
	data := tmplData{
		OpenWeatherCurrent: w,
		Units: units{
			Temp:      "F",
			WindSpeed: "mph",
		},
	}
	funcs := template.FuncMap{
		"Temperature": func(f float32, units string) string { return fmt.Sprintf("%10.1f%s", f, units) },
		"WindSpeed": func(f float32, units string) string {
			num := strconv.FormatFloat(float64(f), 'f', 1, 32)
			return fmt.Sprintf("%11s", num+" "+units)
		},
		"Int": func(n int, units string) string {
			num := strconv.Itoa(n)
			return fmt.Sprintf("%11s", num+units)
		},
	}
	t := template.Must(template.New("").Funcs(funcs).Parse(tmpl))
	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		log.Fatalf("error rendering template: %s", err.Error())
	}

	return buf.String()
}
