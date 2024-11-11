package main

// run command go run ./cmd/web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ngharagazlo/bookings/pkg/config"
	"github.com/ngharagazlo/bookings/pkg/handlers"
	"github.com/ngharagazlo/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8000"

var session *scs.SessionManager
var app config.AppConfig

func main() {

	// get the templateCache from the config

	// change this to true while in production
	app.InProduction = false

	// the (the web page) session lasts for 24 hours
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// the cookies persist after the browser win was closed by the user
	session.Cookie.SameSite = http.SameSiteLaxMode
	// cookies are encrypted (https)
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// Create the tamplates
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("startingthe app on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	// Defining the listening server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
