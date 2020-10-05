# weather-report

Retrieve current weather conditions using the [OpenWeather API](https://openweathermap.org/api).

Weather conditions can be retrieved by
- City name + [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-2:US) state code
- Zip code
- Lat Lon geographical coordinates

# API Key
The [OpenWeather API](https://openweathermap.org/api) requires an API key. [Sign up](https://home.openweathermap.org/users/sign_up) and log in to create a key.

##  There are two ways to supply the API key to the *weather* app:

### Environment variable
Set `WEATHER_API_KEY` to the value of your key:
```
# in shell
export WEATHER_API_KEY=<key>
```
### Config file
If the `WEATHER_API_KEY` environment variable isn't set, *weather* will look for a file called `.weather-api-key` and will load the key from this file.

# Examples
## City name
```
$ weather --searchby=city 'new york,us-ny'
{"coord":{"lon":-74.01,"lat":40.71},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds"}],"base":"stations","main":{"temp":64.85,"feels_like":59.38,"temp_min":63,"temp_max":66.2,"pressure":1021,"humidity":59},"visibility":10000,"wind":{"speed":9.93,"deg":141},"clouds":{"all":90},"dt":1601937557,"sys":{"type":1,"id":5141,"country":"US","sunrise":1601895419,"sunset":1601937081},"timezone":-14400,"id":5128581,"name":"New York","cod":200}
```
## Pretty printed output
```
$ weather --pretty --searchby=city 'new york,us-ny'
New York Weather:
Clouds
Temp             64.8F
Feels like       59.4F
High             66.2F
Low              63.0F
Pressure      1021 hPa 
Humidity           59%
Wind           9.9 mph
Direction  141 degrees
```

## Short help
Run *weather* with no arguments to see a helpful, brief usage message:
```
$ weather
2020/10/05 17:42:43 missing searchby option
usage: weather --searchby <type of term> <term>
example openweather terms:
city: New York, NY
zipcode: 74129
latlon: 38.505252,-90.430133
```

### Author
David P.C. Wollmann  
david.wollmann@gmail.com  
(918) 994-2422
