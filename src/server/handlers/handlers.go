package handlers

import (
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

// TODO: Add comment documentation (type DatabaseHandler)
type DatabaseHandler struct {
	DB     *gorm.DB
	client *redis.Client
}

// TODO: Add comment documentation (type CookieHandler)
type CookieHandler struct {
	store *sessions.CookieStore
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
