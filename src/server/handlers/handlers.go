// This package demonstrates godocs.
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

// TODO: Add comment documentation (type DatabaseHandler)
type DatabaseHandler struct {
	DB     *gorm.DB
	Client *redis.Client
}

// TODO: Add comment documentation (type CookieHandler)
type CookieHandler struct {
	Store *sessions.CookieStore
}

// TODO: Add comment documentation (type AngularHandler)
type AngularHandler struct {
	Host         string
	HTTPAddress  string
	ReverseProxy *httputil.ReverseProxy
}

// TODO: Add comment documentation (func NewDatabaseHandler)
func NewDatabaseHandler(postgresDB *gorm.DB, redisDB *redis.Client) *DatabaseHandler {
	return &DatabaseHandler{postgresDB, redisDB}
}

// TODO: Add comment documentation (func NewCookieHandler)
func NewCookieHandler(cookieStore *sessions.CookieStore) *CookieHandler {
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.Secure = true

	return &CookieHandler{cookieStore}
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
