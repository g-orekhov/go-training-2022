package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/helpers/chi_crud_routes"
	"github.com/test_server/internal/helpers/json_handlers"
)

type UserRoute Route

func NewUserRoute(pattern string, controller interface{}) *UserRoute {
	return &UserRoute{
		pattern:    pattern,
		controller: controller,
	}
}

func (r *UserRoute) Register(apiRouter chi.Router) {
	apiRouter.Group(func(apiRouter chi.Router) {
		apiRouter.Route(r.pattern, func(apiRouter chi.Router) {
			// add CRUD rotes
			chi_crud_routes.AddCrudRoutes(apiRouter, r.controller)
		})
		apiRouter.Handle("/*", json_handlers.NotFoundJSON())
	})
}
