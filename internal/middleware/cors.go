// Package middleware adds any website middleware to http.Handler
package middleware

import "net/http"

const (
	// SiteHeader is a unique HTTP header for the site
	SiteHeader = "colinbagley.dev"
	// SiteValue is the value for the unique site header
	SiteValue = "UP"
)

// Cors adds the CORS HTTP headers
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set(SiteHeader, SiteValue)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
