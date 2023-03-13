package handlers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// TODO: Add comment documentation (type AngularHandler)
type AngularHandler struct {
	Host         string
	HTTPAddress  string
	ReverseProxy *httputil.ReverseProxy
}

// TODO: Add comment documentation (func NewAngularHandler)
func NewAngularHandler(host string, httpAddress string) *AngularHandler {
	origin, err := url.Parse(host)
	if err != nil {
		log.Fatal("Failed to parse frontend network address for origin", err)
	}

	director := func(request *http.Request) {
		request.Header.Add("X-Forwarded-Host", request.Host)
		request.Header.Add("X-Origin-Host", origin.Host)
		request.URL.Scheme = "http"
		request.URL.Host = origin.Host
	}

	reverseProxy := &httputil.ReverseProxy{Director: director}

	return &AngularHandler{host, httpAddress, reverseProxy}
}
