// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOneResourceParams creates a new GetOneResourceParams object
// no default values defined in spec.
func NewGetOneResourceParams() GetOneResourceParams {

	return GetOneResourceParams{}
}

// GetOneResourceParams contains all the bound params for the get one resource operation
// typically these are obtained from a http.Request
//
// swagger:parameters get_one_resource
type GetOneResourceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Resource unique identifier
	  Required: true
	  In: path
	*/
	ResourceID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetOneResourceParams() beforehand.
func (o *GetOneResourceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rResourceID, rhkResourceID, _ := route.Params.GetOK("resource_id")
	if err := o.bindResourceID(rResourceID, rhkResourceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindResourceID binds and validates parameter ResourceID from path.
func (o *GetOneResourceParams) bindResourceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("resource_id", "path", "strfmt.UUID", raw)
	}
	o.ResourceID = *(value.(*strfmt.UUID))

	if err := o.validateResourceID(formats); err != nil {
		return err
	}

	return nil
}

// validateResourceID carries on validations for parameter ResourceID
func (o *GetOneResourceParams) validateResourceID(formats strfmt.Registry) error {

	if err := validate.FormatOf("resource_id", "path", "uuid", o.ResourceID.String(), formats); err != nil {
		return err
	}
	return nil
}
