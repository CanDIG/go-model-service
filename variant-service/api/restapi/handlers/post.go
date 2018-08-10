// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package generics

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"

	"github.com/go-openapi/runtime/middleware"

	"github.com/CanDIG/go-model-service/variant-service/transformations"

	"github.com/CanDIG/go-model-service/variant-service/errors"

	"github.com/CanDIG/go-model-service/variant-service/api/restapi/handlers/utilities"

	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
)

// PostIndividual processes a Individual posted by the API request and creates it into the database.
// It then retrieves the newly created Individual from the database and returns it, along with its URL location.
func PostIndividual(params operations.PostindividualParams) middleware.Responder {
	funcName := "handlers.Post"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewPostindividualInternalServerError().WithPayload(errPayload)
	}

	_, err := getindividualByID(params.individual.ID.String(), tx)
	if err == nil { // TODO this is not a great check
		message := "This Individual already exists in the database. " +
			"It cannot be overwritten with POST; please use PUT instead."
		errors.Log(nil, 405, funcName, message)
		errPayload := &apimodels.Error{Code: 405001, Message: &message}
		return operations.NewPostindividualMethodNotAllowed().WithPayload(errPayload)
	}

	newindividual, errPayload := transformations.individualAPIToDataModel(*params.individual, tx)
	if errPayload != nil {
		return operations.NewPostindividualInternalServerError().WithPayload(errPayload)
	}

	err = tx.Create(newindividual)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Create into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostindividualInternalServerError().WithPayload(errPayload)
	}

	retreivedDataindividual, err := getindividualByID(newindividual.ID.String(), tx)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Failed to get Individual by ID from database immediately following its creation")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostindividualInternalServerError().WithPayload(errPayload)
	}

	retreivedAPIindividual, errPayload := transformations.individualDataToAPIModel(*retreivedDataindividual)
	if err != nil {
		return operations.NewPostindividualInternalServerError().WithPayload(errPayload)
	}

	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + retreivedAPIindividual.ID.String()
	return operations.NewPostindividualCreated().WithPayload(retreivedAPIindividual).WithLocation(location)
}

// PostVariant processes a Variant posted by the API request and creates it into the database.
// It then retrieves the newly created Variant from the database and returns it, along with its URL location.
func PostVariant(params operations.PostvariantParams) middleware.Responder {
	funcName := "handlers.Post"

	tx, errPayload := utilities.ConnectDevelopment(funcName)
	if errPayload != nil {
		return operations.NewPostvariantInternalServerError().WithPayload(errPayload)
	}

	_, err := getvariantByID(params.variant.ID.String(), tx)
	if err == nil { // TODO this is not a great check
		message := "This Variant already exists in the database. " +
			"It cannot be overwritten with POST; please use PUT instead."
		errors.Log(nil, 405, funcName, message)
		errPayload := &apimodels.Error{Code: 405001, Message: &message}
		return operations.NewPostvariantMethodNotAllowed().WithPayload(errPayload)
	}

	newvariant, errPayload := transformations.variantAPIToDataModel(*params.variant, tx)
	if errPayload != nil {
		return operations.NewPostvariantInternalServerError().WithPayload(errPayload)
	}

	err = tx.Create(newvariant)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Create into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostvariantInternalServerError().WithPayload(errPayload)
	}

	retreivedDatavariant, err := getvariantByID(newvariant.ID.String(), tx)
	if err != nil {
		errors.Log(err, 500, funcName,
			"Failed to get Variant by ID from database immediately following its creation")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostvariantInternalServerError().WithPayload(errPayload)
	}

	retreivedAPIvariant, errPayload := transformations.variantDataToAPIModel(*retreivedDatavariant)
	if err != nil {
		return operations.NewPostvariantInternalServerError().WithPayload(errPayload)
	}

	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + retreivedAPIvariant.ID.String()
	return operations.NewPostvariantCreated().WithPayload(retreivedAPIvariant).WithLocation(location)
}
