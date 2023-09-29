package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {

	mux := chi.NewMux()

	mux.Use(noSurf)
	mux.Use(middleware.Logger)
	mux.Get("/", app.home)
	mux.Get("/about", app.about)

	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	return mux

}
