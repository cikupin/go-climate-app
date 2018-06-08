package interfaces

import (
	"github.com/cikupin/go-climate-app/models"
)

// IAjaxService - interface for AjaxService
type IAjaxRepository interface {
	GetOpenWeatherData(city string, days int) models.OpenWeatherResponse
}
