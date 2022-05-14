package routes

import "github.com/go-chi/chi/v5"

type UserRoute Route

func NewUserRoute(pattern string, controller interface{}) *UserRoute {
	return &UserRoute{
		pattern:    pattern,
		controller: controller,
	}
}

func (r *UserRoute) Register(apiRouter chi.Router) {
	apiRouter.Group(func(apiRouter chi.Router) {
		AddCrudRoutes(&apiRouter, r.controller, r.pattern)
		apiRouter.Handle("/*", NotFoundJSON())
	})
}
