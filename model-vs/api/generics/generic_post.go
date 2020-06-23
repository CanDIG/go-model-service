package generics

import (
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/utilities"
	"github.com/CanDIG/go-model-service/model-vs/errors"
	"github.com/CanDIG/go-model-service/tools/log"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// PostIndividual processes a Individual posted by the API request and creates it into the database.
// It then retrieves the newly created Individual from the database and returns it, along with its URL location.
func PostIndividual(params operations.PostIndividualParams) middleware.Responder {
	tx, errPayload := utilities.ConnectDevelopment(params.HTTPRequest)
	if errPayload != nil {
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	_, err := utilities.GetIndividualByID(params.Individual.ID.String(), tx)
	if err == nil { // TODO this is not a great check
		message := "This Individual already exists in the database. " +
			"It cannot be overwritten with POST; please use PUT instead."
		code := 405001

		log.Write(params.HTTPRequest, code, nil).Warn(message)
		errPayload := &apimodels.Error{Code: int64(code), Message: &message}
		return operations.NewPostIndividualMethodNotAllowed().WithPayload(errPayload)
	}

	newIndividual, errPayload := individualAPIToDataModel(*params.Individual, params.HTTPRequest, tx)
	if errPayload != nil {
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	err = tx.Create(newIndividual)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Create into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	// TODO if errors occur from this point on, the Individual may have already been created,
	// so it should be deleted prior to return
	retrievedDataIndividual, err := utilities.GetIndividualByID(newIndividual.ID.String(), tx)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Failed to get Individual by ID from database immediately following its creation")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	retrievedAPIIndividual, errPayload := individualDataToAPIModel(*retrievedDataIndividual, params.HTTPRequest)
	if errPayload != nil {
		return operations.NewPostIndividualInternalServerError().WithPayload(errPayload)
	}

	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + retrievedAPIIndividual.ID.String()
	return operations.NewPostIndividualCreated().WithPayload(retrievedAPIIndividual).WithLocation(location)
}
