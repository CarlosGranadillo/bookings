package main

import (
	"net/http"

	"github.com/CarlosGranadillo/bookings/pkg/config"
	"github.com/CarlosGranadillo/bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(appConfigPointer *config.ApplicationConfig) http.Handler {
	// mux := pat.New()
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsoleOnPageViewed)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// create a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	// use the file server
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
