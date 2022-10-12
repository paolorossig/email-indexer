package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/paolorossig/go-challenge/adapter/zincsearch"
	"github.com/paolorossig/go-challenge/handler"
	"github.com/paolorossig/go-challenge/service"
)

const (
	defaultPort = "4000"
)

// App describes the application.
type App struct {
	Server       *chi.Mux
	httpClient   *http.Client
	dependencies *dependencies
}

// dependencies has the data for the application dependency injection
type dependencies struct {
	emailHandler   *handler.EmailHandler
	indexerHandler *handler.IndexerHandler
}

func (a *App) setupInfrastructure() {
	a.httpClient = &http.Client{}
}

func (a *App) setupDependencies() {
	emailService := service.NewEmailService()
	emailHandler := handler.NewEmailHandler(emailService)

	zincSearchAdapter := zincsearch.NewClient(a.httpClient)
	indexerService := service.NewIndexerService(zincSearchAdapter)
	indexerHandler := handler.NewIndexerHandler(indexerService, emailService)

	a.dependencies = &dependencies{
		emailHandler:   emailHandler,
		indexerHandler: indexerHandler,
	}
}

func (a *App) setupServer() {
	a.Server = chi.NewRouter()
	a.Server.Use(middleware.RequestID)
	a.Server.Use(middleware.Logger)
	a.Server.Use(middleware.Recoverer)
	a.Server.Use(middleware.URLFormat)
	a.Server.Use(render.SetContentType(render.ContentTypeJSON))

	a.SetupRoutes()
}

// NewApp loads the infrastructure and dependencies of the app.
func NewApp() *App {
	a := &App{}
	a.setupInfrastructure()
	a.setupDependencies()
	a.setupServer()

	return a
}

// StartApp loads the application with its routes.
func (a *App) StartApp() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	fmt.Printf("App is running on: http://localhost:%s ...\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), a.Server); err != nil {
		panic(err)
	}
}
