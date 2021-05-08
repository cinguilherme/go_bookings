package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/cinguilherme/playground/pkg/config"
	"github.com/cinguilherme/playground/pkg/handlers"
	"github.com/cinguilherme/playground/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// must be true in production
	app.InProduction = false

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("unable to create template cache")
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	app.TemplateCache = tc
	app.UserCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	serv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
