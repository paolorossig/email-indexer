package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/paolorossig/email-indexer/pkg/adapter/zincsearch"
	"github.com/paolorossig/email-indexer/pkg/handler"
	"github.com/paolorossig/email-indexer/pkg/service"
)

const (
	defaultPort    = "8000"
	defaultAppName = "email-indexer-api"
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
	zincSearchAdapter := zincsearch.NewClient(a.httpClient)

	emailService := service.NewEmailService(zincSearchAdapter)
	emailHandler := handler.NewEmailHandler(emailService)

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
	appName := os.Getenv("APP_NAME")
	port := os.Getenv("PORT")

	if port == "" || appName == "" {
		port = defaultPort
		appName = defaultAppName
		log.Println("Env variables APP_NAME and PORT are not set. Using default values.")
	}

	log.Printf("App %s is running on: http://localhost:%s\n", appName, port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), a.Server); err != nil {
		panic(err)
	}
}
