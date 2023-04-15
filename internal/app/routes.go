package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *App) Routes() http.Handler {
	router := chi.NewRouter()

	// set custom error handler for 404 Not Found response
	router.NotFound(http.HandlerFunc(app.NotFoundResponse))

	// set custom error handler for 405 Method Not Allowed response
	router.MethodNotAllowed(http.HandlerFunc(app.MethodNotAllowedResponse))

	// CORS allowing calls from any domain
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/healthcheck", app.HealthcheckHandler)
	router.Get("/image/{resolution}", app.ImageHandler)

	return router
}
