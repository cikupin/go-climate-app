package routes

import (
	"github.com/cikupin/go-climate-app/controllers"
	"github.com/cikupin/go-climate-app/repositories"
	"github.com/cikupin/go-climate-app/services"
)

// DependencyInjector - stuct for DependencyInjector
type DependencyInjector struct {
}

// Inject dependency into Router
func (di *DependencyInjector) Inject(r *Router) {
	// Repositories
	ajaxRepository := new(repositories.AjaxRepository)

	// Services
	ajaxService := new(services.AjaxService)
	ajaxService.AjaxRepository = ajaxRepository

	// Controllers
	ajaxController := new(controllers.AjaxController)
	ajaxController.AjaxService = ajaxService
	indexController := new(controllers.IndexController)

	// Routes
	r.AjaxController = ajaxController
	r.IndexController = indexController
}
