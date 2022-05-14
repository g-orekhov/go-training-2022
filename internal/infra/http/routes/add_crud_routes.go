package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// TODO: uncomment if using go v1.18
// type CRUD interface {
// 	getRoutes | createRoute | updateRoute | deleteRoute
// }

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

// TODO controller must be CRUD interface (for the go v1.18)
/*
Generates standard CRUD routes for the controller if it impliments CRUD interface
*/
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
				"/{id}",
				updateController.Update(),
			)
		}
		deleteController, ok := controller.(deleteRoute)
		if ok {
			apiRouter.Delete(
				"/{id}",
				deleteController.Delete(),
			)
		}
	})
}
