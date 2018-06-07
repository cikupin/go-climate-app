package interfaces

import (
	"net/http"
)

// IAjaxController - interface for AjaxController
type IAjaxController interface {
	GetWeather(w http.ResponseWriter, r *http.Request)
}
