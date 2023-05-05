package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// -----
	// routing using pat
	// -------
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// ------
	// routing using chi
	// -------
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // absorbs panic and prints stack trace

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
