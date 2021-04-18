// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteBeneficiaryByIDHandlerFunc turns a function with the right signature into a delete beneficiary by Id handler
type DeleteBeneficiaryByIDHandlerFunc func(DeleteBeneficiaryByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBeneficiaryByIDHandlerFunc) Handle(params DeleteBeneficiaryByIDParams) middleware.Responder {
	return fn(params)
}

// DeleteBeneficiaryByIDHandler interface for that can handle valid delete beneficiary by Id params
type DeleteBeneficiaryByIDHandler interface {
	Handle(DeleteBeneficiaryByIDParams) middleware.Responder
}

// NewDeleteBeneficiaryByID creates a new http.Handler for the delete beneficiary by Id operation
func NewDeleteBeneficiaryByID(ctx *middleware.Context, handler DeleteBeneficiaryByIDHandler) *DeleteBeneficiaryByID {
	return &DeleteBeneficiaryByID{Context: ctx, Handler: handler}
}

/*DeleteBeneficiaryByID swagger:route DELETE /registration/beneficiaries/{id} deleteBeneficiaryById

delete a beneficiary by id

*/
type DeleteBeneficiaryByID struct {
	Context *middleware.Context
	Handler DeleteBeneficiaryByIDHandler
}

func (o *DeleteBeneficiaryByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBeneficiaryByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}