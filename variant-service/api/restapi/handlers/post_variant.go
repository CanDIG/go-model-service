package handlers

import (
	"github.com/CanDIG/go-model-service/variant-service/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/CanDIG/go-model-service/variant-service/transformations"
	"github.com/gobuffalo/pop"
	apimodels "github.com/CanDIG/go-model-service/variant-service/api/models"
	"github.com/CanDIG/go-model-service/variant-service/errors"
)

// PostVariants processes a variant posted by the API request and creates it into the database.
// It then retrieves the newly created variant from the database and returns it, along with its URL location.
func PostVariants(params operations.PostVariantParams) middleware.Responder {
	tx, err := pop.Connect("development")
	if err != nil {
		errors.Log(err, 500,"restapi.api.MainPostVariantHandler",
			"Failed to connect to database: development")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}

	_, err = getVariantByID(params.Variant.ID.String(), tx)
	if err == nil { // TODO this is not a great check
		message := "This variant already exists in the database. " +
			"It cannot be overwritten with POST; please use PUT instead."
		errors.Log(nil, 405,"restapi.api.MainPostVariantHandler", message)
		errPayload := &apimodels.Error{Code: 405001, Message: &message}
		return operations.NewPostVariantMethodNotAllowed().WithPayload(errPayload)
	}

	newVariant, errPayload := transformations.VariantAPIToDataModel(*params.Variant, tx)
	if errPayload != nil {
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}

	_, err = tx.ValidateAndCreate(newVariant)
	if err != nil {
		errors.Log(err, 500,"restapi.api.MainPostVariantHandler",
			"ValidateAndCreate into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}

	dataVariant, err := getVariantByID(newVariant.ID.String(), tx)
	if err != nil {
		errors.Log(err, 500,"restapi.api.MainPostVariantHandler, restapi.getVariantByID(string)",
			"Failed to get variant by ID from database immediately following its creation")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}

	apiVariant, errPayload := transformations.VariantDataToAPIModel(*dataVariant)
	if err != nil {
		return operations.NewPostVariantInternalServerError().WithPayload(errPayload)
	}
	
	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + apiVariant.ID.String()
	return operations.NewPostVariantCreated().WithPayload(apiVariant).WithLocation(location)
}