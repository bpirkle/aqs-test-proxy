package handlers

import (
	"net/http"

	"gitlab.wikimedia.org/frankie/aqsassist"
)

// NotFoundHandler is the HTTP handler when no match routes are found.

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	problemResp := aqsassist.CreateProblem(http.StatusNotFound, "Invalid route", r.URL.RequestURI())
	(*problemResp).WriteTo(w)
}
