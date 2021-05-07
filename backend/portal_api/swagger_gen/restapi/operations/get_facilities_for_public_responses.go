// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetFacilitiesForPublicOKCode is the HTTP code returned for type GetFacilitiesForPublicOK
const GetFacilitiesForPublicOKCode int = 200

/*GetFacilitiesForPublicOK OK

swagger:response getFacilitiesForPublicOK
*/
type GetFacilitiesForPublicOK struct {

	/*
	  In: Body
	*/
	Payload *GetFacilitiesForPublicOKBody `json:"body,omitempty"`
}

// NewGetFacilitiesForPublicOK creates GetFacilitiesForPublicOK with default headers values
func NewGetFacilitiesForPublicOK() *GetFacilitiesForPublicOK {

	return &GetFacilitiesForPublicOK{}
}

// WithPayload adds the payload to the get facilities for public o k response
func (o *GetFacilitiesForPublicOK) WithPayload(payload *GetFacilitiesForPublicOKBody) *GetFacilitiesForPublicOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get facilities for public o k response
func (o *GetFacilitiesForPublicOK) SetPayload(payload *GetFacilitiesForPublicOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFacilitiesForPublicOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
