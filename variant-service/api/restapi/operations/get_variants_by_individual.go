// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetVariantsByIndividualHandlerFunc turns a function with the right signature into a get variants by individual handler
type GetVariantsByIndividualHandlerFunc func(GetVariantsByIndividualParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVariantsByIndividualHandlerFunc) Handle(params GetVariantsByIndividualParams) middleware.Responder {
	return fn(params)
}

// GetVariantsByIndividualHandler interface for that can handle valid get variants by individual params
type GetVariantsByIndividualHandler interface {
	Handle(GetVariantsByIndividualParams) middleware.Responder
}

// NewGetVariantsByIndividual creates a new http.Handler for the get variants by individual operation
func NewGetVariantsByIndividual(ctx *middleware.Context, handler GetVariantsByIndividualHandler) *GetVariantsByIndividual {
	return &GetVariantsByIndividual{Context: ctx, Handler: handler}
}

/*GetVariantsByIndividual swagger:route GET /individuals/{individual_id}/variants getVariantsByIndividual

Get variants called in an individual

*/
type GetVariantsByIndividual struct {
	Context *middleware.Context
	Handler GetVariantsByIndividualHandler
}

func (o *GetVariantsByIndividual) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetVariantsByIndividualParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
