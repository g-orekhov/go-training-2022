package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/test_server/internal/helpers/json_handlers"
)

type routeInterface interface {
	Register(chi.Router)
}

func Router(routes ...routeInterface) http.Handler {
	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", json_handlers.NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			for _, route := range routes {
				route.Register(apiRouter)
			}

			apiRouter.Handle("/*", json_handlers.NotFoundJSON())
		})
	})

	return router
}
