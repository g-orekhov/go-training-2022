package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/test_server/internal/infra/http/controllers"
)

func Router(eventController *controllers.EventController) http.Handler {
	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/v1", func(apiRouter chi.Router) {

			apiRouter.Group(func(apiRouter chi.Router) {
				AddCrudRoutes(&apiRouter, eventController, "/events")

				apiRouter.Handle("/*", NotFoundJSON())
			})
			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	return router
}