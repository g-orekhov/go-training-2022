package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type getRoutes interface {
	FindAll() http.HandlerFunc
	FindOne() http.HandlerFunc
}

type createRoute interface {
	Create() http.HandlerFunc
}

type updateRoute interface {
	Update() http.HandlerFunc
}

type deleteRoute interface {
	Delete() http.HandlerFunc
}

func AddCrudRoutes(router *chi.Router, controller interface{}, pattern string) {
	(*router).Route(pattern, func(apiRouter chi.Router) {
		getController, ok := controller.(getRoutes)
		if ok {
			apiRouter.Get(
				"/",
				getController.FindAll(),
			)
			apiRouter.Get(
				"/{id}",
				getController.FindOne(),
			)
		}
		createController, ok := controller.(createRoute)
		if ok {
			apiRouter.Post(
				"/",
				createController.Create(),
			)
		}
		updateController, ok := controller.(updateRoute)
		if ok {
			apiRouter.Put(
				"/",
				updateController.Update(),
			)
		}
		deleteController, ok := controller.(deleteRoute)
		if ok {
			apiRouter.Delete(
				"/",
				deleteController.Delete(),
			)
		}
	})
}
