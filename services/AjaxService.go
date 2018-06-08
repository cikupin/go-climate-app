package services

import (
	"time"

	"github.com/cikupin/go-climate-app/interfaces"
	"github.com/cikupin/go-climate-app/libraries"
	"github.com/cikupin/go-climate-app/models"
	"github.com/cikupin/go-climate-app/repositories"
)

// DailyTemperature struct
type DailyTemperature struct {
	Date    string  `json:"date"`
	DayTemp float64 `json:"daytemp"`
	VarTemp float64 `json:"vartemp"`
}

// AjaxService - struct for AjaxService
type AjaxService struct {
	AjaxRepository interfaces.IAjaxRepository
}

// GetCityWeather call AjaxRepository.GetOpenWeatherData() to obtain weather from openweathermap.org
func (r *AjaxService) GetCityWeather(city string) models.WeatherResponse {
	r.AjaxRepository = new(repositories.AjaxRepository)
	var cityWeather models.WeatherResponse

	openWeatherData := r.AjaxRepository.GetOpenWeatherData(city)
	cityWeather.City = openWeatherData.City.Name

	var dailyTemp DailyTemperature
	var totalTemp, totalVar float64
	for _, v := range openWeatherData.List {
		dailyTemp.Date = time.Unix(v.DayTime, 0).Format("2006-01-02")
		dailyTemp.DayTemp = libraries.RoundNumber(v.Temp.Day)
		dailyTemp.VarTemp = libraries.RoundNumber(v.Temp.Max - v.Temp.Min)

		totalTemp += dailyTemp.DayTemp
		totalVar += dailyTemp.VarTemp
		cityWeather.Temp = append(cityWeather.Temp, dailyTemp)
	}

	cityWeather.Average.Temp = libraries.RoundNumber(totalTemp / 5)
	cityWeather.Average.Var = libraries.RoundNumber(totalVar / 5)

	return cityWeather
}
