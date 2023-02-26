// Package handlers contains HTTP handlers for the application.
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"server/config"
	"server/models"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

/*
Interface for accessing the DB functions in models.User. Is used to mock DB functions in unit tests of HTTP handler functions.
*/
type UserService interface {
	CreateUser(user *models.User) (uint64, error)
	FindUser(userID uint64) (*models.User, error)
	FindUserByEmail(userEmail string) (*models.User, error)
	UpdateUser(userID uint64, user *models.User) (*models.User, error)
	DeleteUser(userID uint64) (bool, error)
}

/*
TODO: migrate existing handler functions using the Handler struct to Env struct so that they can be unit tested
*/
type Handler struct {
	DB          *gorm.DB
	redisClient *redis.Client
}

/*
Env is a struct that holds the necessary dependencies for handling HTTP requests related to user management.

Fields:
- user: Implements the UserService interface which provides access to the database for user management operations and allows for mocking of the Env struct to unit test HTTP handler functions.
- store: A *sessions.CookieStore object that provides session management operations using cookies.
*/
type Env struct {
	users UserService
	store *sessions.CookieStore
}

// TODO: Add comment documentation (type AngularHandler)
type AngularHandler struct {
	Host         string
	HTTPAddress  string
	ReverseProxy *httputil.ReverseProxy
}

/*
NewHandler creates a new instance of the Handler struct with the provided dependencies and returns a pointer to it.
The Handler struct contains a *gorm.DB, *redis.Client, and *sessions.CookieStore as attributes.
The cookieStore's HttpOnly and Secure options are set to true.

Parameters:
- postgresDB: A pointer to a *gorm.DB instance representing the PostgreSQL database connection.
- redisDB: A pointer to a *redis.Client instance representing the Redis database connection.
- cookieStore: A pointer to a *sessions.CookieStore instance representing the session cookie store.

Returns:
- A pointer to a new Handler instance with the provided dependencies.
*/
func NewHandler(postgresDB *gorm.DB, redisDB *redis.Client) *Handler {
	return &Handler{postgresDB, redisDB}
}

/*
NewEnv returns an instance of the Env struct with provided dependencies.
*/
func NewEnv(postgresDB *gorm.DB, cookieStore *sessions.CookieStore) *Env {
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.Secure = true

	user := &models.UserEnv{DB: postgresDB}

	return &Env{users: user, store: cookieStore}
}

// TODO: Add comment documentation (func NewAngularHandler)
func NewAngularHandler() *AngularHandler {
	var host string = config.AppConfig.FRONTEND_HOST
	var httpAddress = fmt.Sprintf("http://%s", config.AppConfig.GetFrontendNetworkAddress())

	origin, err := url.Parse(host)
	if err != nil {
		log.Fatal("Failed to parse frontend network address for origin", err)
	}

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
	}

	reverseProxy := &httputil.ReverseProxy{Director: director}

	return &AngularHandler{host, httpAddress, reverseProxy}
}
