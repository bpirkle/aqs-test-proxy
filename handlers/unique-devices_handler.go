package handlers

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gerrit.wikimedia.org/r/mediawiki/services/servicelib-golang/logger"
)

// UniqueDevicesHandler is the HTTP handler for unique-devices endpoint requests.
type UniqueDevicesHandler struct {
	Logger *logger.Logger
}

func (s *UniqueDevicesHandler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	reqUrl := r.URL.RequestURI()
	reqLogger := s.Logger.Request(r)

	getUrl := "https://staging.svc.eqiad.wmnet:4972" + reqUrl

	t := &http.Transport{
		TLSHandshakeTimeout: 600 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   90 * 30 * time.Second,
		Transport: t,
	}

	res, err := client.Get(getUrl)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	reqLogger.Log(logger.INFO, string(resBody))

	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
