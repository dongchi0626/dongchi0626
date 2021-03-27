// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InitializeFacilitySlotsHandlerFunc turns a function with the right signature into a initialize facility slots handler
type InitializeFacilitySlotsHandlerFunc func(InitializeFacilitySlotsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn InitializeFacilitySlotsHandlerFunc) Handle(params InitializeFacilitySlotsParams) middleware.Responder {
	return fn(params)
}

// InitializeFacilitySlotsHandler interface for that can handle valid initialize facility slots params
type InitializeFacilitySlotsHandler interface {
	Handle(InitializeFacilitySlotsParams) middleware.Responder
}

// NewInitializeFacilitySlots creates a new http.Handler for the initialize facility slots operation
func NewInitializeFacilitySlots(ctx *middleware.Context, handler InitializeFacilitySlotsHandler) *InitializeFacilitySlots {
	return &InitializeFacilitySlots{Context: ctx, Handler: handler}
}

/* InitializeFacilitySlots swagger:route POST /facility/slots/init initializeFacilitySlots

Initialize facility slots

*/
type InitializeFacilitySlots struct {
	Context *middleware.Context
	Handler InitializeFacilitySlotsHandler
}

func (o *InitializeFacilitySlots) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewInitializeFacilitySlotsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// InitializeFacilitySlotsBody initialize facility slots body
//
// swagger:model InitializeFacilitySlotsBody
type InitializeFacilitySlotsBody struct {

	// api key
	APIKey string `json:"apiKey,omitempty"`
}

// Validate validates this initialize facility slots body
func (o *InitializeFacilitySlotsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this initialize facility slots body based on context it is used
func (o *InitializeFacilitySlotsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InitializeFacilitySlotsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InitializeFacilitySlotsBody) UnmarshalBinary(b []byte) error {
	var res InitializeFacilitySlotsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}