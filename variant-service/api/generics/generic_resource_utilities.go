package generics

import (
	"github.com/gobuffalo/pop"
		datamodels "github.com/CanDIG/go-model-service/variant-service/data/models"
)

// getResourceByID returns the Resource in the database corresponding to the given ID (or nil if no match is found)
func getResourceByID(id string, tx *pop.Connection) (*datamodels.Resource, error) {
	var individual *datamodels.Resource
	err := tx.Find(individual, id)
	return individual, err
}