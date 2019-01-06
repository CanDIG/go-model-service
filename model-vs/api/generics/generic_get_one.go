package generics

import (
	"github.com/CanDIG/go-model-service/tools/log"
	"github.com/sirupsen/logrus"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/utilities"
	"github.com/CanDIG/go-model-service/model-vs/errors"
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
)

// GetOneIndividual returns the Individual in the database that corresponds to a given UUID.
func GetOneIndividual(params operations.GetOneIndividualParams) middleware.Responder {
	tx, errPayload := utilities.ConnectDevelopment(params.HTTPRequest)
	if errPayload != nil {
		return operations.NewGetOneIndividualInternalServerError().WithPayload(errPayload)
	}

	dataIndividual, err := utilities.GetIndividualByID(params.IndividualID.String(), tx)
	if err != nil {
		message := "This Individual cannot be found."
		code := 404001
		
		log.Write(params.HTTPRequest, code, err).Warn(message)
		errPayload := &apimodels.Error{Code: int64(code), Message: &message}
		return operations.NewGetOneIndividualNotFound().WithPayload(errPayload)
	}

	apiIndividual, errPayload := individualDataToAPIModel(*dataIndividual, params.HTTPRequest)
	if errPayload != nil {
		return operations.NewGetOneIndividualInternalServerError().WithPayload(errPayload)
	}

	return operations.NewGetOneIndividualOK().WithPayload(apiIndividual)
}