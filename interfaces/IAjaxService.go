package interfaces

import (
	"github.com/cikupin/go-climate-app/models"
)

// IAjaxService - interface for AjaxService
type IAjaxService interface {
	GetCityWeather(city string, days int) models.WeatherResponse
}
