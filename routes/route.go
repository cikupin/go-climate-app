package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
}

func (r *Router) Handler() http.Handler {
	routes := mux.NewRouter()

	return routes
}
