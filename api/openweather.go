package api

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
