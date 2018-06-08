package routes

import (
	"net/http"

	"github.com/cikupin/go-climate-app/controllers"
	"github.com/cikupin/go-climate-app/interfaces"
	"github.com/gorilla/mux"
)

// Router - struct for Router
type Router struct {
	IndexController interfaces.IIndexController
	AjaxController  interfaces.IAjaxController
}

// Handler handling application routing
func (r *Router) Handler() *mux.Router {
	r.IndexController = new(controllers.IndexController)
	r.AjaxController = new(controllers.AjaxController)

	fs := http.FileServer(http.Dir("public"))
	routes := mux.NewRouter()

	routes.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	routes.HandleFunc("/", r.IndexController.Index).Methods("GET")
	routes.HandleFunc("/get-weather/{city}/{days}", r.AjaxController.GetWeather).Methods("GET")

	return routes
}
