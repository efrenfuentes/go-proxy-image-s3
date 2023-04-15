package app

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type App struct {
	Config Config
	Logger *log.Logger
}

func NewApp(config Config, log *log.Logger) *App {
	return &App{
		Config: config,
		Logger: log,
	}
}

func (app *App) Run() error {
	server := &http.Server{
		Addr:         app.Config.HttpAddr,
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Infof("Starting server on %s", server.Addr)
	return server.ListenAndServe()
}
