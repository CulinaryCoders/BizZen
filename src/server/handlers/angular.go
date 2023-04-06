package handlers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

/*
*Description*

type AngularHandler

The AngularHandler type is used to store/access various attributes of the frontend Angular application
*/
type AngularHandler struct {
	Host         string                 // The host address for the frontend Angular application
	HTTPAddress  string                 // The host address for the frontend Angular application in http format
	ReverseProxy *httputil.ReverseProxy // The reverse proxy used to allow the frontend to connect to the backend server
}

/*
*Description*

func NewAngularHandler

Creates a new AngularHandler instance.

*Parameters*

	host  <string>

	   The host address of the Angular application (only the host portion of the address, no port number).

	httpAddress  <string>

		The http address of the Angular application in 'http://host:port' format.

*Returns*

	_  <AngularHandler>

		A pointer to the new AngularHandler instance that was created.
*/
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
