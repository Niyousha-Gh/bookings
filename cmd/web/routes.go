package main

import (
	"net/http"

	"github.com/ngharagazlo/bookings/pkg/config"
	"github.com/ngharagazlo/bookings/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// There is the template cache in the config
func routes(app *config.AppConfig) http.Handler {

	// this is to use the downloaded http package to have Get (in go.mod)
	//mux := pat.New()

	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	// a chi middleware to handle the panics instead of klling the app
	// loading the packages installed in the go.mod
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux

}
