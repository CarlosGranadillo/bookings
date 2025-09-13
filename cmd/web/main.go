package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CarlosGranadillo/bookings/pkg/config"
	"github.com/CarlosGranadillo/bookings/pkg/handlers"
	"github.com/CarlosGranadillo/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var appConfig config.ApplicationConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// change to true when in production and using https
	appConfig.IsProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // survive browser restarts
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.IsProd

	appConfig.Session = session

	// create the template cache
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error trying to create the template cache")
	}

	appConfig.TemplateCache = templateCache
	appConfig.UseCache = true

	// Give access to the app config (to the render package)
	render.InitAppConfig(&appConfig)

	// Handlers repository
	handlersRepo := handlers.NewRepo(&appConfig)
	handlers.InitHandlersRepo(handlersRepo)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfig),
	}

	fmt.Printf("Starting application on port %s\n", portNumber)

	// Start the server
	err = srv.ListenAndServe()

	// srv.ListenAndServe() on success -> blocks forever (is a blocking call)
	// the error handling code after srv.ListenAndServe() only executes if the server fails to start
	// That is why it's not strictly necessary to add the line below inside a -> if err != nil {}
	// Still might not hurt to be explicit, but in my understanding it doesn't make any difference
	log.Fatal(err)
}
