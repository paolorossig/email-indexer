package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// SetupRoutes sets up the routes for the app.
func (a *App) SetupRoutes() {
	a.Server.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "Welcome to the Go Challenge!"})
	})
	a.Server.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
	})
	a.Server.Mount("/indexer", indexerRouter(a))
	a.Server.Mount("/emails", emailsRouter(a))
}

func indexerRouter(a *App) chi.Router {
	r := chi.NewRouter()
	r.Post("/emails", a.dependencies.indexerHandler.IndexEmails)

	return r
}

func emailsRouter(a *App) chi.Router {
	r := chi.NewRouter()
	r.Get("/search", a.dependencies.emailHandler.SearchInEmails)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", a.dependencies.emailHandler.GetAvailableUsers)
		r.Get("/{userID}", a.dependencies.emailHandler.GetEmailsFromUser)
	})

	return r
}
