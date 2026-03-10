package config

import (
	h "ascii-web/internal/handlers"
	p "ascii-web/internal/processor"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router  *http.ServeMux
	Skecher h.AsciiSketch
}

func (a *App) InitializeRouters() {
	a.Router.HandleFunc("GET /", a.Skecher.HandleVerbs)
	a.Router.HandleFunc("POST /ascii-art", a.Skecher.HandleVerbs)
}

func (a *App) InitializeServer() {
	processor := p.CreateNewProcessor()
	sketcher := h.CreateNewAsciiSketcher(processor)

	a.Skecher = *sketcher
	a.Router = http.NewServeMux()
	a.InitializeRouters()
}

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
