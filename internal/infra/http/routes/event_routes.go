package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/helpers/chi_crud_routes"
	"github.com/test_server/internal/helpers/json_handlers"
)

type EventRoute Route

type EventController interface {
	FindAll() http.HandlerFunc
	FindOne() http.HandlerFunc
	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
	GetNearby() http.HandlerFunc
}

func NewEventRoute(pattern string, controller EventController) *EventRoute {
	return &EventRoute{
		pattern:    pattern,
		controller: controller,
	}
}

func (r *EventRoute) Register(apiRouter chi.Router) {
	apiRouter.Group(func(apiRouter chi.Router) {
		chi_crud_routes.AddCrudRoutes(&apiRouter, r.controller, r.pattern)
		controller, _ := r.controller.(EventController)
		apiRouter.Get(r.pattern+"/geo/{lat},{long},{dist}", controller.GetNearby())
		apiRouter.Handle("/*", json_handlers.NotFoundJSON())
	})
}
