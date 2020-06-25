package utilities

import (
	"fmt"
	apimodels "github.com/CanDIG/go-model-service/model-vs/api/models"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/CanDIG/go-model-service/utilities/log"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gobuffalo/pop"
)

// addAND only adds an AND to the given conditions string if it already has contents.
func addAND(conditions string) string {
	if conditions == "" {
		return ""
	}
	return conditions + " AND "
}

// GetIndividuals builds an Individuals-specific query out of the given parameters.
// Since there are presently no parameters expected for this request, it simply returns all Individuals.
func GetIndividuals(params operations.GetIndividualsParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	return pop.Q(tx), nil
}

// GetCalls builds an Calls-specific query out of the given parameters.
// Since there are presently no parameters expected for this request, it simply returns all Calls.
func GetCalls(params operations.GetCallsParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	return pop.Q(tx), nil
}

// GetVariants builds an Individuals-specific query out of the given parameters.
// It rejects get-all requests, as such a request would, in a production service, return a prohibitively
// large amount of data and would likely only be entered in error or in malice.
// May return a 403: Forbidden response.
func GetVariants(params operations.GetVariantsParams, tx *pop.Connection) (*pop.Query, middleware.Responder) {
	conditions := ""

	if params.Chromosome != nil {
		conditions = fmt.Sprintf(addAND(conditions)+"chromosome = '%s'", *params.Chromosome)
	}
	if params.Start != nil {
		conditions = fmt.Sprintf(addAND(conditions)+"start >= '%d'", *params.Start)
	}
	if params.End != nil {
		conditions = fmt.Sprintf(addAND(conditions)+"start <= '%d'", *params.End)
	}

	if conditions == "" {
		message := "Forbidden to query for all variants. " +
			"Please provide parameters in the query string for 'chromosome', 'start', and/or 'end'."
		code := 403001

		log.Write(params.HTTPRequest, code, nil).Warn(message)
		errPayload := &apimodels.Error{Code: int64(code), Message: &message}
		return nil, operations.NewGetVariantsForbidden().WithPayload(errPayload)
	}

	query := tx.Where(conditions)
	return query, nil
}
