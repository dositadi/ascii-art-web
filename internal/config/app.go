package config

import (
	h "ascii-web/internal/handlers"
	p "ascii-web/internal/processor"
	"log"
	"net/http"
	"time"
)

// The app struct where the system begins to function
type App struct {
	Router  *http.ServeMux
	Skecher h.AsciiSketch
}

// This function initializes all routes in  the software
func (a *App) InitializeRouters() {
	a.Router.HandleFunc("GET /", a.Skecher.HandleVerbs)
	a.Router.HandleFunc("POST /ascii-art", a.Skecher.HandleVerbs)
}

// This function initializes the server
func (a *App) InitializeServer() {
	processor := p.CreateNewProcessor()
	sketcher := h.CreateNewAsciiSketcher(processor)

	a.Skecher = *sketcher
	a.Router = http.NewServeMux()
	a.InitializeRouters()
}

// This is where the app begins to run
func (a *App) Run() {
	a.InitializeServer()

	server := http.Server{
		Addr:         ":8081",
		Handler:      a.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
