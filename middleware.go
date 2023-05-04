package main

import (
	"net/http"
)

// Set content type as application/json
func SetContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

// SecureHeadersMiddleware adds two basic security headers to each HTTP response
// X-XSS-Protection: 1; mode-block can help to prevent XSS attacks
// X-Frame-Options: deny can help to prevent clickjacking attacks
func SecureHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode-block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}
