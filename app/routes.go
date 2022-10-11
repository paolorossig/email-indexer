package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// SetupRoutes sets up the routes for the app.
func (a *App) SetupRoutes() {
	a.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "Welcome to the Go Challenge!"})
	})
	a.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
	})
	a.Router.Mount("/indexer", indexerRouter(a))
}

func indexerRouter(a *App) chi.Router {
	r := chi.NewRouter()
	r.Get("/", a.dependencies.indexerHandler.Index)
	r.Get("/available", a.dependencies.indexerHandler.Index)
	return r
}
