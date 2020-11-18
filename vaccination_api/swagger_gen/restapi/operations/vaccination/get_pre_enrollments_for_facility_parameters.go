// Code generated by go-swagger; DO NOT EDIT.

package vaccination

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetPreEnrollmentsForFacilityParams creates a new GetPreEnrollmentsForFacilityParams object
// no default values defined in spec.
func NewGetPreEnrollmentsForFacilityParams() GetPreEnrollmentsForFacilityParams {

	return GetPreEnrollmentsForFacilityParams{}
}

// GetPreEnrollmentsForFacilityParams contains all the bound params for the get pre enrollments for facility operation
// typically these are obtained from a http.Request
//
// swagger:parameters getPreEnrollmentsForFacility
type GetPreEnrollmentsForFacilityParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	FacilityCode string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetPreEnrollmentsForFacilityParams() beforehand.
func (o *GetPreEnrollmentsForFacilityParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rFacilityCode, rhkFacilityCode, _ := route.Params.GetOK("facilityCode")
	if err := o.bindFacilityCode(rFacilityCode, rhkFacilityCode, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFacilityCode binds and validates parameter FacilityCode from path.
func (o *GetPreEnrollmentsForFacilityParams) bindFacilityCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.FacilityCode = raw

	return nil
}