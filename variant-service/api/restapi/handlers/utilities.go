package handlers

import (
	"github.com/gobuffalo/pop"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	"github.com/CanDIG/go-model-service/variant-service/errors"
)

// connectDevelopment connects to the development database and returns the connection and/or error message
func connectDevelopment(funcName string) (*pop.Connection, *apimodels.Error) {
	tx, err := pop.Connect("development")
	if err != nil {
		errors.Log(err, 500, funcName, "Failed to connect to database: development")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}
	return tx, nil
}