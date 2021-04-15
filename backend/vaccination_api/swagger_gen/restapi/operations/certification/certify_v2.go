// Code generated by go-swagger; DO NOT EDIT.

package certification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/divoc/api/swagger_gen/models"
)

// CertifyV2HandlerFunc turns a function with the right signature into a certify v2 handler
type CertifyV2HandlerFunc func(CertifyV2Params, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn CertifyV2HandlerFunc) Handle(params CertifyV2Params, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// CertifyV2Handler interface for that can handle valid certify v2 params
type CertifyV2Handler interface {
	Handle(CertifyV2Params, *models.JWTClaimBody) middleware.Responder
}

// NewCertifyV2 creates a new http.Handler for the certify v2 operation
func NewCertifyV2(ctx *middleware.Context, handler CertifyV2Handler) *CertifyV2 {
	return &CertifyV2{Context: ctx, Handler: handler}
}

/*CertifyV2 swagger:route POST /v2/certify certification certifyV2

Certify the one or more vaccination

Certification happens asynchronously, this requires vaccinator athorization and vaccinator should be trained for the vaccination that is being certified.

*/
type CertifyV2 struct {
	Context *middleware.Context
	Handler CertifyV2Handler
}

func (o *CertifyV2) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCertifyV2Params()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CertifyV2BadRequestBody certify v2 bad request body
//
// swagger:model CertifyV2BadRequestBody
type CertifyV2BadRequestBody struct {

	// code
	// Required: true
	Code *string `json:"code"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this certify v2 bad request body
func (o *CertifyV2BadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CertifyV2BadRequestBody) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("certifyV2BadRequest"+"."+"code", "body", o.Code); err != nil {
		return err
	}

	return nil
}

func (o *CertifyV2BadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("certifyV2BadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CertifyV2BadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CertifyV2BadRequestBody) UnmarshalBinary(b []byte) error {
	var res CertifyV2BadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}