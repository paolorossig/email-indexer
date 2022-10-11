package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/paolorossig/go-challenge/handler"
	"github.com/paolorossig/go-challenge/service"
)

// App describes the application.
type App struct {
	Router       *chi.Mux
	dependencies *dependencies
}

// dependencies has the data for the application dependency injection
type dependencies struct {
	indexerHandler *handler.IndexerHandler
}

func (a *App) setupDependencies() {
	a.dependencies = &dependencies{}
	indexerService := service.NewIndexerService()
	indexerHandler := handler.NewIndexerHandler(indexerService)
	a.dependencies.indexerHandler = indexerHandler
}

func (a *App) setupServer() {
	a.Router = chi.NewRouter()
	a.Router.Use(middleware.RequestID)
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(middleware.URLFormat)
	a.Router.Use(render.SetContentType(render.ContentTypeJSON))

	a.SetupRoutes()
}

// NewApp loads the infrastructure and dependencies of the app.
func NewApp() *App {
	a := &App{}
	a.setupDependencies()
	a.setupServer()

	return a
}

// StartApp loads the application with its routes.
func (a *App) StartApp() {
	if err := http.ListenAndServe(":4000", a.Router); err != nil {
		panic(err)
	}
}
