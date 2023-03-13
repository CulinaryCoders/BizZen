package handlers

import (
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/database"
	"server/middleware"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

var App *Application = &Application{}

// TODO:  Add documentation (type Application)
type Application struct {
	Router      *mux.Router
	CookieStore *sessions.CookieStore
	AppDB       *gorm.DB
	CacheDB     *redis.Client
	NGHandler   *AngularHandler
}

// TODO:  Add documentation (func Initialize)
func (app *Application) Initialize(appDBName string) {
	// Initialize main app database
	var dbConnectionString string = config.AppConfig.GetPostgresDBConnectionString(appDBName)
	var debug bool = config.AppConfig.DEBUG_MODE
	App.AppDB = database.InitializePostgresDB(dbConnectionString, debug)

	// Initialize cache db
	var cacheDSN string = config.AppConfig.GetRedisDBNetworkAddress()
	App.CacheDB = database.InitializeRedisDB(cacheDSN)

	// Initialize cookie store
	App.CookieStore = sessions.NewCookieStore([]byte("super-secret"))
	App.CookieStore.Options.HttpOnly = true
	App.CookieStore.Options.Secure = true

	// Initialize AngularHandler
	var ngHost string = config.AppConfig.FRONTEND_HOST
	var ngHttpAddress string = fmt.Sprintf("http://%s", config.AppConfig.GetFrontendNetworkAddress())
	app.NGHandler = NewAngularHandler(ngHost, ngHttpAddress)

	// Initialize router and routes
	app.Router = mux.NewRouter()
	app.Router.Use(middleware.RequestLoggingMiddleware)
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
	app.Router.HandleFunc("/register", app.CreateUser).Methods("POST")
	app.Router.HandleFunc("/authenticate", app.Authenticate).Methods("POST")
	app.Router.HandleFunc("/user/{id}", app.GetUser).Methods("GET")
	app.Router.HandleFunc("/user/{id}", app.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/user/{id}", app.DeleteUser).Methods("DELETE")

	// Business routes
	app.Router.HandleFunc("/business", app.CreateBusiness).Methods("POST")
	app.Router.HandleFunc("/business/{id}", app.GetBusiness).Methods("GET")
	app.Router.HandleFunc("/business/{id}", app.UpdateBusiness).Methods("PUT")
	app.Router.HandleFunc("/business/{id}", app.DeleteBusiness).Methods("DELETE")

	// Business routes
	app.Router.HandleFunc("/address", app.CreateAddress).Methods("POST")
	app.Router.HandleFunc("/address/{id}", app.GetAddress).Methods("GET")
	app.Router.HandleFunc("/address/{id}", app.UpdateAddress).Methods("PUT")
	app.Router.HandleFunc("/address/{id}", app.DeleteAddress).Methods("DELETE")

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
