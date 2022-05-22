package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/helpers/chi_crud_routes"
	"github.com/test_server/internal/helpers/json_handlers"
)

type EventRoute Route

func NewEventRoute(pattern string, controller interface{}) *EventRoute {
	return &EventRoute{
		pattern:    pattern,
		controller: controller,
	}
}

func (r *EventRoute) Register(apiRouter chi.Router) {
	apiRouter.Group(func(apiRouter chi.Router) {
		chi_crud_routes.AddCrudRoutes(&apiRouter, r.controller, r.pattern)
		apiRouter.Handle("/*", json_handlers.NotFoundJSON())
	})
}
