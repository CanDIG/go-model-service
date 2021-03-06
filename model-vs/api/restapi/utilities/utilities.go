/*
Package utilities implements general-purpose utility functions for use by the restapi handlers.
*/
package utilities

import (
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	"github.com/CanDIG/go-model-service/model-vs/errors"
	"github.com/CanDIG/go-model-service/tools/log"
	"github.com/gobuffalo/pop"
	"net/http"
)

// ConnectDevelopment connects to the development database and returns the connection and/or error message
func ConnectDevelopment(HTTPRequest *http.Request) (*pop.Connection, *apimodels.Error) {
	tx, err := pop.Connect("development")
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("Failed to connect to database: development")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}
	return tx, nil
}
