package repositories

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cikupin/go-climate-app/models"
)

// AjaxRepository - struct for AjaxRepository
type AjaxRepository struct {
}

// GetOpenWeatherData obtain weather data from openweathermap.org
func (r *AjaxRepository) GetOpenWeatherData(city string, days int) models.OpenWeatherResponse {
	var bufferURL bytes.Buffer
	bufferURL.WriteString("http://api.openweathermap.org/data/2.5/forecast/daily?q=")
	bufferURL.WriteString(city)
	bufferURL.WriteString("&mode=json&units=metric&cnt=")
	bufferURL.WriteString(strconv.Itoa(days))
	bufferURL.WriteString("&APPID=481e3bc28e5264e5607c2b65b449bfc1")

	res, err := http.Get(bufferURL.String())
	if err != nil {
		log.Fatal(err)
	}

	var w models.OpenWeatherResponse
	err = json.NewDecoder(res.Body).Decode(&w)
	if err != nil {
		log.Fatal(err)
	}

	return w
}
