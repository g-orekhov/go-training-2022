package routes

import "github.com/go-chi/chi/v5"

type EventRoute Route

func NewEventRoute(pattern string, controller interface{}) *EventRoute {
	return &EventRoute{
		pattern:    pattern,
		controller: controller,
	}
}

func (r *EventRoute) Register(apiRouter chi.Router) {
	apiRouter.Group(func(apiRouter chi.Router) {
		AddCrudRoutes(&apiRouter, r.controller, r.pattern)
		apiRouter.Handle("/*", NotFoundJSON())
	})
}
