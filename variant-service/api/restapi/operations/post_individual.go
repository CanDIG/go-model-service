// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostIndividualHandlerFunc turns a function with the right signature into a post individual handler
type PostIndividualHandlerFunc func(PostIndividualParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostIndividualHandlerFunc) Handle(params PostIndividualParams) middleware.Responder {
	return fn(params)
}

// PostIndividualHandler interface for that can handle valid post individual params
type PostIndividualHandler interface {
	Handle(PostIndividualParams) middleware.Responder
}

// NewPostIndividual creates a new http.Handler for the post individual operation
func NewPostIndividual(ctx *middleware.Context, handler PostIndividualHandler) *PostIndividual {
	return &PostIndividual{Context: ctx, Handler: handler}
}

/*PostIndividual swagger:route POST /individuals postIndividual

Add an individual to the database

*/
type PostIndividual struct {
	Context *middleware.Context
	Handler PostIndividualHandler
}

func (o *PostIndividual) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostIndividualParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}