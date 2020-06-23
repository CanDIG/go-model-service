/*
Package log implements logrus-powered logging functionality
*/
package log

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// TODO should we somehow configure api.logger in configure_variant_service.go ?

// TODO make all of this initialization configurable
// Init initializes the logger according to command-line-provided config parameters.
func Init() {
	// Log as JSON instead of the default ASCII formatter
	logrus.SetFormatter(&logrus.JSONFormatter{})

	//	log.SetOutput()

	logrus.SetLevel(logrus.WarnLevel)
}

// TODO log pid
// Write employs logrus to produce a logger in a consistent format.
// This logger extracts pertinent HTTP request information from the net/http Request parameter
// provided in go-swagger auto-generated restapi/operations *_parameters.go files.
func Write(HTTPRequest *http.Request, httpCode int, err error) *logrus.Entry {
	entry := logrus.WithFields(logrus.Fields{
		"service": "go-model-service",
		"version": "0.0.1",
		"host":    HTTPRequest.Host,
		"ip":      HTTPRequest.RemoteAddr})

	if httpCode != 0 {
		entry = entry.WithField("httpCode", strconv.Itoa(httpCode))
	}

	if err != nil {
		entry = entry.WithField("error", err.Error())
	}

	return entry
}
