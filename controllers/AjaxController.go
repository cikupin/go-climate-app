package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cikupin/go-climate-app/interfaces"
	"github.com/cikupin/go-climate-app/services"
	"github.com/gorilla/mux"
)

// AjaxController - struct for AjaxController
type AjaxController struct {
	AjaxService interfaces.IAjaxService
}

// GetWeather call AjaxService.GetCityWeather() to get weather of certain city
func (a *AjaxController) GetWeather(w http.ResponseWriter, r *http.Request) {
	a.AjaxService = new(services.AjaxService)
	vars := mux.Vars(r)
	city := vars["city"]
	days, err := strconv.Atoi(vars["days"])

	weather := a.AjaxService.GetCityWeather(city, days)
	j, err := json.Marshal(weather)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
