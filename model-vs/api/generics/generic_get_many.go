package generics

import (
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/utilities"
	datamodels "github.com/CanDIG/go-model-service/model-vs/data/models"
	"github.com/CanDIG/go-model-service/model-vs/errors"
	"github.com/CanDIG/go-model-service/tools/log"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// GetIndividuals returns all Individuals in the database given zero or more query parameters.
// The query parameters are handled separately in getIndividualsQuery.
func GetIndividuals(params operations.GetIndividualsParams) middleware.Responder {
	tx, errPayload := utilities.ConnectDevelopment(params.HTTPRequest)
	if errPayload != nil {
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	// the full error response is handled here rather than the payload because a variety of http codes may occur
	query, errResponse := utilities.GetIndividuals(params, tx)
	if errResponse != nil {
		return errResponse
	}

	var dataIndividuals []datamodels.Individual
	err := query.All(&dataIndividuals)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Problems getting Individuals from database")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetIndividualsInternalServerError().WithPayload(errPayload)
	}

	var apiIndividuals []*apimodels.Individual
	for _, dataIndividual := range dataIndividuals {
		apiIndividual, errPayload := individualDataToAPIModel(dataIndividual, params.HTTPRequest)
		if errPayload != nil {
			return operations.NewGetIndividualsInternalServerError().WithPayload(errPayload)
		}
		apiIndividuals = append(apiIndividuals, apiIndividual)
	}

	return operations.NewGetIndividualsOK().WithPayload(apiIndividuals)
}
