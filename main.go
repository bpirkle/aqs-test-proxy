package main

import (
	"device-analytics/handlers"
	"fmt"
	"net/http"
	"os"
	"path"

	log "gerrit.wikimedia.org/r/mediawiki/services/servicelib-golang/logger"
	"github.com/gorilla/mux"
)

var (
	// These values are assigned at build using `-ldflags` (see: Makefile)
	buildDate = "unknown"
	buildHost = "unknown"
	version   = "unknown"
)

// Entrypoint for the service
func main() {
	var err error
	serviceName := "aqs-test-proxy"
	baseURI := "/metrics/unique-devices"
	logger, err := log.NewLogger(os.Stdout, serviceName, "DEBUG")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to initialize the logger: %s", err)
		os.Exit(1)
	}

	logger.Info("Initializing service %s (Go version: %s, Build host: %s, Timestamp: %s", serviceName, version, buildHost, buildDate)

	// pass bound struct method to handler
	uniqueDevicesHandler := &handlers.UniqueDevicesHandler{
		Logger: logger}
	healthz := NewHealthz(version, buildDate, buildHost)

	r := mux.NewRouter().SkipClean(true).UseEncodedPath()
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

	r.HandleFunc("/healthz", healthz.HandleHTTP).Methods("GET")

	r.Use(SetContentType, SecureHeadersMiddleware)

	r.HandleFunc(path.Join(baseURI, "/{project}/{access-site}/{granularity}/{start}/{end}"), uniqueDevicesHandler.HandleHTTP).Methods("GET")

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "localhost", 8086),
		Handler: r,
	}

	err = srv.ListenAndServe()
	fmt.Println(err.Error())
}
