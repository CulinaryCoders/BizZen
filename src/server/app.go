package main

import (
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/database"
	"server/handlers"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
)

// TODO:  Add documentation (type Application)
type Application struct {
	Router    *mux.Router
	Handler   *handlers.Handler
	Env       *handlers.Env
	NGHandler *handlers.AngularHandler
}

// TODO:  Add documentation (func Initialize)
func (app *Application) Initialize(appDBName string) {
	// Initialize main app database
	var dbConnectionString string = config.AppConfig.GetPostgresDBConnectionString(appDBName)
	appDB := database.InitializePostgresDB(dbConnectionString)

	// Initialize cache db
	var cacheDSN string = config.AppConfig.GetRedisDBNetworkAddress()
	cacheDB := database.InitializeRedisDB(cacheDSN)

	// Initialize cookie store and CookieHandler
	cookieStore := sessions.NewCookieStore([]byte("super-secret"))

	// Initialize Handler
	app.Handler = handlers.NewHandler(appDB, cacheDB)

	//Initialize Env for HTTP handlers
	app.Env = handlers.NewEnv(appDB, cookieStore)

	// Initialize AngularHandler
	app.NGHandler = handlers.NewAngularHandler()

	// Initialize router and routes
	app.Router = mux.NewRouter()
	app.Router.Use(handlers.RequestLoggingMiddleware)
	app.initializeRoutes()
}

// TODO:  Add documentation (func initializeRoutes)
func (app *Application) initializeRoutes() {
	// Define routes
	app.Router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// TODO: Create API table of contents or list of valid objects in API for root "/" route
		fmt.Fprint(writer, "Hello, World!")
	})

	// User routes
	app.Router.HandleFunc("/register", app.Env.RegisterUser).Methods("POST")
	app.Router.HandleFunc("/authenticate", app.Env.Authenticate).Methods("POST")
	app.Router.HandleFunc("/user/{id}", app.Env.GetUser).Methods("GET")
	app.Router.HandleFunc("/user/{id}", app.Env.UpdateUser).Methods("POST")
	app.Router.HandleFunc("/user/{id}", app.Env.DeleteUser).Methods("DELETE")

	// Business routes
	app.Router.HandleFunc("/business", app.Handler.CreateBusiness).Methods("POST")
	app.Router.HandleFunc("/business/{id}", app.Handler.GetBusiness).Methods("GET")
	app.Router.HandleFunc("/business/{id}", app.Handler.UpdateBusiness).Methods("PUT")
	app.Router.HandleFunc("/business/{id}", app.Handler.DeleteBusiness).Methods("DELETE")

	// Path prefix for API to work with Angular frontend
	// WARNING: This MUST be the last route defined by the router.
	app.Router.PathPrefix("/").Handler(app.NGHandler.ReverseProxy).Methods("GET")
}

// TODO:  Add documentation (func Run)
func (app *Application) Run(networkAddress string) {
	var appHTTPAddress string = fmt.Sprintf("http://%s", networkAddress)

	corsOptions := cors.Options{
		AllowedOrigins:      []string{appHTTPAddress, networkAddress, app.NGHandler.HTTPAddress, app.NGHandler.Host},
		AllowedMethods:      []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders:      []string{"X-Requested-With", "Content-Type", "Authorization", "DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Range", "Range"},
		ExposedHeaders:      []string{"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Content-Range", "Range", "Content-Disposition"},
		MaxAge:              86400,
		AllowCredentials:    true,
		AllowPrivateNetwork: true,
	}

	httpHandler := cors.New(corsOptions).Handler(app.Router)

	log.Printf("Server is now listening on: %s", networkAddress)
	log.Fatal(http.ListenAndServe(networkAddress, httpHandler))
}
