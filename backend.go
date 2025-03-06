package gobackend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cpgoodwi/gobackend/config"
	"github.com/cpgoodwi/gobackend/db"
)

type Backend struct {
	Config  config.AppConfig
	DB      *sql.DB
	Router  *http.ServeMux
	Logger  *log.Logger
	Started time.Time
}

func NewBackend() (*Backend, error) {
	// default config
	config := config.AppConfig{
		DataSource: "postgresql://root:password@localhost:5433/gobackend?sslmode=disable",
		DataDir:    "/data",
		ServerURL:  "http://localhost",
		Port:       8080,
	}

	// TODO: load config from file

	// ensure data directory exists
	if err := os.MkdirAll(config.DataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	// intialize database
	db, err := db.GetDB()
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	app := &Backend{
		Config:  config,
		DB:      db,
		Router:  http.NewServeMux(),
		Logger:  log.New(os.Stdout, "[Backend]", log.LstdFlags),
		Started: time.Now(),
	}

	return app, nil
}

func (app *Backend) Start() error {
	// register routes
	app.registerRoutes()

	// start server
	serverAddr := fmt.Sprintf("%s:%d", app.Config.ServerURL, app.Config.Port)
	app.Logger.Printf("starting server on %s", serverAddr)

	return http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Port), app.Router)
}

func (app *Backend) registerRoutes() {
	// api routes

	// admin ui routes
}
