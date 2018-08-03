// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/CanDIG/go-model-service/variant-service/api/models"
)

// GetCallsOKCode is the HTTP code returned for type GetCallsOK
const GetCallsOKCode int = 200

/*GetCallsOK Return calls

swagger:response getCallsOK
*/
type GetCallsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Call `json:"body,omitempty"`
}

// NewGetCallsOK creates GetCallsOK with default headers values
func NewGetCallsOK() *GetCallsOK {

	return &GetCallsOK{}
}

// WithPayload adds the payload to the get calls o k response
func (o *GetCallsOK) WithPayload(payload []*models.Call) *GetCallsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get calls o k response
func (o *GetCallsOK) SetPayload(payload []*models.Call) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCallsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Call, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
