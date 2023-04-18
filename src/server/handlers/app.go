package handlers

import (
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/middleware"
	"server/models"

	//"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

var App *Application = &Application{}

/*
*Description*

type Application

Application is intended to be used as a singleton object for storing/accessing global application state
*/
type Application struct {
	Router      *mux.Router           // Gorilla Mux Router (used to define/configure API endpoints)
	CookieStore *sessions.CookieStore // Gorilla Sessions CookieStore for storing session/cookie data
	AppDB       *gorm.DB              // gorm.DB instance used as main application database
	// CacheDB     *redis.Client         // redis.Client instance used for caching database
	NGHandler *AngularHandler // AngularHandler that allows the frontend to connect to the backend API server
}

/*
*Description*

func Initialize

Initializes the various components of the Application instance (router, database(s), cookies/session, handlers, etc.)

*Parameters*

	appDBName  <string>

		The name of the database that will be used as the main application database.

*Returns*

	None
*/
func (app *Application) Initialize(appDBName string) {
	// Initialize main app database
	var dbConnectionString string = config.AppConfig.GetPostgresDBConnectionString(appDBName)
	App.AppDB = models.InitializePostgresDB(dbConnectionString, config.Debug)

	// Initialize cache db
	// var cacheDSN string = config.AppConfig.GetRedisDBNetworkAddress()
	// App.CacheDB = models.InitializeRedisDB(cacheDSN)

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

/*
*Description*

func initializeRoutes

Defines the the API endpoints/routes for the application and their behavior using the application's Gorilla Mux Router.

*Parameters*

	None

*Returns*

	None
*/
func (app *Application) initializeRoutes() {
	// Define routes
	app.Router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// TODO: Create API table of contents or list of valid objects in API for root "/" route
		fmt.Fprint(writer, "Hello, World!")
	})

	// User routes
	app.Router.HandleFunc("/register", app.CreateUser).Methods("POST")
	app.Router.HandleFunc("/login", app.Authenticate).Methods("POST")
	app.Router.HandleFunc("/user/{id}", app.GetUser).Methods("GET")
	app.Router.HandleFunc("/user/{id}", app.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/user/{id}", app.DeleteUser).Methods("DELETE")
	app.Router.HandleFunc("/users", app.GetUsers).Methods("GET")
	app.Router.HandleFunc("/user/{id}/service-appointments", app.GetUserServiceAppointments).Methods("GET")

	// Business routes
	app.Router.HandleFunc("/business", app.CreateBusiness).Methods("POST")
	app.Router.HandleFunc("/business/{id}", app.GetBusiness).Methods("GET")
	app.Router.HandleFunc("/business/{id}", app.UpdateBusiness).Methods("PUT")
	app.Router.HandleFunc("/business/{id}", app.DeleteBusiness).Methods("DELETE")
	app.Router.HandleFunc("/businesses", app.GetBusinesses).Methods("GET")
	app.Router.HandleFunc("/business/{id}/services", app.GetBusinessServices).Methods("GET")
	app.Router.HandleFunc("/business/{id}/service-appointments", app.GetBusinessServiceAppointments).Methods("GET")

	// Address routes
	// app.Router.HandleFunc("/address", app.CreateAddress).Methods("POST")
	// app.Router.HandleFunc("/address/{id}", app.GetAddress).Methods("GET")
	// app.Router.HandleFunc("/address/{id}", app.UpdateAddress).Methods("PUT")
	// app.Router.HandleFunc("/address/{id}", app.DeleteAddress).Methods("DELETE")

	// Service routes
	app.Router.HandleFunc("/service", app.CreateService).Methods("POST")
	app.Router.HandleFunc("/service/{id}", app.GetService).Methods("GET")
	app.Router.HandleFunc("/service/{id}", app.UpdateService).Methods("PUT")
	app.Router.HandleFunc("/service/{id}", app.DeleteService).Methods("DELETE")
	app.Router.HandleFunc("/services", app.GetServices).Methods("GET")
	app.Router.HandleFunc("/service/{service-id}/user/{user-id}", app.GetUserEnrolledStatus).Methods("GET")
	app.Router.HandleFunc("/service/{id}/users", app.GetListOfEnrolledUsers).Methods("GET")
	app.Router.HandleFunc("/service/{id}/user-count", app.GetEnrolledUsersCount).Methods("GET")
	app.Router.HandleFunc("/service/{id}/appointments", app.GetActiveServiceAppointments).Methods("GET")
	app.Router.HandleFunc("/service/{id}/appointments/active", app.GetActiveServiceAppointments).Methods("GET")
	app.Router.HandleFunc("/service/{id}/appointments/all", app.GetServiceAppointments).Methods("GET")
	// TODO: app.Router.HandleFunc("/service/{id}/user-appointments", app.GetUserAppointments).Methods("GET")

	// Appointment routes
	app.Router.HandleFunc("/appointment", app.CreateAppointment).Methods("POST")
	app.Router.HandleFunc("/appointment/{id}", app.GetAppointment).Methods("GET")
	app.Router.HandleFunc("/appointment/{id}", app.UpdateAppointment).Methods("PUT")
	app.Router.HandleFunc("/appointment/{id}", app.DeleteAppointment).Methods("DELETE")
	app.Router.HandleFunc("/appointments", app.GetActiveAppointments).Methods("GET")
	app.Router.HandleFunc("/appointments/active", app.GetActiveAppointments).Methods("GET")
	app.Router.HandleFunc("/appointments/all", app.GetAppointments).Methods("GET")
	app.Router.HandleFunc("/appointment/{id}/cancel", app.CancelAppointment).Methods("POST")

	// Invoice routes
	app.Router.HandleFunc("/invoice", app.CreateInvoice).Methods("POST")
	app.Router.HandleFunc("/invoice/{id}", app.GetInvoice).Methods("GET")
	app.Router.HandleFunc("/invoice/{id}", app.UpdateInvoice).Methods("PUT")
	app.Router.HandleFunc("/invoice/{id}", app.DeleteInvoice).Methods("DELETE")
	app.Router.HandleFunc("/invoices", app.GetInvoices).Methods("GET")

	// Path prefix for API to work with Angular frontend
	// WARNING: This MUST be the last route defined by the router.
	app.Router.PathPrefix("/").Handler(app.NGHandler.ReverseProxy).Methods("GET")
}

/*
*Description*

func Run

Serves the application on the specified network address.

*Parameters*

	networkAddress <string>

	   The network address to use for the application server in 'host:port' format.

*Returns*

	None
*/
func (app *Application) Run(networkAddress string) {
	var appHTTPAddress string = fmt.Sprintf("http://%s", networkAddress)

	corsOptions := cors.Options{
		AllowedOrigins:      []string{appHTTPAddress, networkAddress, app.NGHandler.HTTPAddress, app.NGHandler.Host},
		AllowedMethods:      []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
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
