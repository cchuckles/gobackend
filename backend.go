package gobackend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cchuckles/gobackend/config"
	"github.com/cchuckles/gobackend/db"
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

	be := &Backend{
		Config:  config,
		DB:      db,
		Router:  http.NewServeMux(),
		Logger:  log.New(os.Stdout, "[Backend]", log.LstdFlags),
		Started: time.Now(),
	}

	return be, nil
}

func (be *Backend) Start() error {
	// register routes
	be.registerRoutes()

	// start server
	serverAddr := fmt.Sprintf("%s:%d", be.Config.ServerURL, be.Config.Port)
	be.Logger.Printf("starting server on %s", serverAddr)

	return http.ListenAndServe(fmt.Sprintf(":%d", be.Config.Port), be.Router)
}

func (be *Backend) registerRoutes() {
	// api routes

	// admin ui routes
}
