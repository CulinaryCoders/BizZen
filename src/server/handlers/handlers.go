// Package handlers contains HTTP handlers for the application.
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"server/config"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

/*
Handler is a struct that holds the necessary dependencies for handling HTTP requests related to user management.

Fields:
- DB: A *gorm.DB object that provides access to the database for user management operations.
- redisClient: A *redis.Client object that provides access to the Redis server for session management operations.
- cookieStore: A *sessions.CookieStore object that provides session management operations using cookies.

Description:
The Handler struct is used to handle HTTP requests related to user management, such as authentication and authorization.
It holds dependencies such as a *gorm.DB object, which provides access to the database for user management operations,
a *redis.Client object, which provides access to the Redis server for session management operations, and a
*sessions.CookieStore object, which provides session management using Gorilla sessions.
*/
type Handler struct {
	DB          *gorm.DB
	redisClient *redis.Client
	cookieStore *sessions.CookieStore
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
func NewHandler(postgresDB *gorm.DB, redisDB *redis.Client, cookieStore *sessions.CookieStore) *Handler {
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.Secure = true
	return &Handler{postgresDB, redisDB, cookieStore}
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
