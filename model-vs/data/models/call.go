package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"time"
)

// Call : 	The ORM-side representation of the Call data object.
// 			A single variant call on a single individual.
// 			Contains some unique data, but is essentially the many-to-many association between variants and individuals.
type Call struct {
	ID           uuid.UUID   `json:"id" db:"id"`
	CreatedAt    time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" db:"updated_at"`
	IndividualID uuid.UUID   `json:"individual_id" db:"individual_id"`
	Individual   *Individual `json:"individual" belongs_to:"individual"`
	VariantID    uuid.UUID   `json:"variant_id" db:"variant_id"`
	Variant      *Variant    `json:"variant" belongs_to:"variant"`
	Genotype     string      `json:"genotype" db:"genotype"`
	Format       string      `json:"format" db:"format"`
}

// String is not required by pop and may be deleted
func (c Call) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Calls is not required by pop and may be deleted
type Calls []Call

// String is not required by pop and may be deleted
func (c Calls) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Call) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: c.IndividualID, Name: "IndividualID"},
		&validators.UUIDIsPresent{Field: c.VariantID, Name: "VariantID"},
		&validators.StringIsPresent{Field: c.Genotype, Name: "Genotype"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Call) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Call) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
