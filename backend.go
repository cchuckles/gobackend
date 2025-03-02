package gobackend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	DatabaseType string `json:"databaseType"`
	ServerURL    string `json:"serverUrl"`
	Port         int    `json:"port"`
}

// TODO: set up configs for other database types

type Backend struct {
	Config  Config
	DB      *sql.DB
	Router  *http.ServeMux
	Logger  *log.Logger
	Started time.Time
}

func NewBackend() (*Backend, error) {
	// default config
	config := Config{
		DatabaseType: "sqlite",
		ServerURL:    "http://localhost",
		Port:         8080,
	}

	// TODO: load config from file

	// intialize database
	var db *sql.DB
	var err error
	switch config.DatabaseType {
	case "sqlite":
		// ensure data directory exists
		if err = os.MkdirAll("./data", 0755); err != nil {
			return nil, fmt.Errorf("failed to create data directory: %w", err)
		}

		dbPath := filepath.Join("./data", "data.db")
		db, err = sql.Open("sqlite3", dbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open database: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.DatabaseType)
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
