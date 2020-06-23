package generics

import (
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/utilities"
	"github.com/CanDIG/go-model-service/model-vs/errors"
	"github.com/CanDIG/go-model-service/tools/log"
	"github.com/go-openapi/runtime/middleware"
)

// GetIndividualsByVariant returns all Individuals with a given Variant called.
// Since Individuals and Variants have a many-to-many relationship, Calls are used as the relation/junction between them.
func GetIndividualsByVariant(params operations.GetIndividualsByVariantParams) middleware.Responder {
	tx, errPayload := utilities.ConnectDevelopment(params.HTTPRequest)
	if errPayload != nil {
		return operations.NewGetIndividualsByVariantInternalServerError().WithPayload(errPayload)
	}

	dataVariant, err := utilities.GetVariantByID(params.VariantID.String(), tx)
	if err != nil {
		message := "The Variant by which you are trying to query by cannot be found."
		code := 404002

		log.Write(params.HTTPRequest, code, err).Warn(message)
		errPayload := &apimodels.Error{Code: int64(code), Message: &message}
		return operations.NewGetIndividualsByVariantNotFound().WithPayload(errPayload)
	}

	err = tx.Load(dataVariant, "Individuals")
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Problems loading individuals from variant in database")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetIndividualsByVariantInternalServerError().WithPayload(errPayload)
	}

	var apiIndividuals []*apimodels.Individual
	for _, dataIndividual := range dataVariant.Individuals {
		apiIndividual, errPayload := individualDataToAPIModel(dataIndividual, params.HTTPRequest)
		if errPayload != nil {
			return operations.NewGetIndividualsByVariantInternalServerError().WithPayload(errPayload)
		}
		apiIndividuals = append(apiIndividuals, apiIndividual)
	}

	return operations.NewGetIndividualsByVariantOK().WithPayload(apiIndividuals)
}
