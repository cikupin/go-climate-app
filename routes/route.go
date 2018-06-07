package routes

import (
	"net/http"

	"github.com/cikupin/go-climate-app/controllers"
	"github.com/cikupin/go-climate-app/interfaces"
	"github.com/gorilla/mux"
)

type Router struct {
	IndexController interfaces.IIndexController
	AjaxController  interfaces.IAjaxController
}

func (r *Router) Handler() http.Handler {
	r.IndexController = new(controllers.IndexController)
	r.AjaxController = new(controllers.AjaxController)

	routes := mux.NewRouter().StrictSlash(false)
	routes.HandleFunc("/", r.IndexController.Index).Methods("GET")
	routes.HandleFunc("/get-weather/{city}", r.AjaxController.GetWeather).Methods("GET")

	return routes
}
