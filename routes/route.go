package routes

import (
	"net/http"

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
	injector := new(DependencyInjector)
	injector.Inject(r)

	fs := http.FileServer(http.Dir("public"))
	routes := mux.NewRouter()

	routes.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	routes.HandleFunc("/", r.IndexController.Index).Methods("GET")
	routes.HandleFunc("/get-weather/{city}/{days}", r.AjaxController.GetWeather).Methods("GET")

	return routes
}
