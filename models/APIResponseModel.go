package models

// OpenWeatherResponse defines reponse from openweather
type OpenWeatherResponse struct {
	City struct {
		Name string `json:"name"`
	} `json:"city"`
	List []struct {
		DayTime int64 `json:"dt"`
		Temp    struct {
			Day float64 `json:"day"`
			Min float64 `json:"min"`
			Max float64 `json:"max"`
		} `json:"temp"`
	} `json:"list"`
}

// WeatherResponse defines weather response to be send back to client
type WeatherResponse struct {
	City string `json:"name"`
	Temp []struct {
		Date    string  `json:"date"`
		DayTemp float64 `json:"daytemp"`
		VarTemp float64 `json:"vartemp"`
	} `json:"temp"`
	Average struct {
		Temp float64 `json:"avgtemp"`
		Var  float64 `json:"avgvar"`
	} `json:"average"`
}
