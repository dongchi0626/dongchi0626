// Code generated by go-swagger; DO NOT EDIT.

package vaccination

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/divoc/api/swagger_gen/models"
)

// GetRecipientsOKCode is the HTTP code returned for type GetRecipientsOK
const GetRecipientsOKCode int = 200

/*GetRecipientsOK OK

swagger:response getRecipientsOK
*/
type GetRecipientsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Recipient `json:"body,omitempty"`
}

// NewGetRecipientsOK creates GetRecipientsOK with default headers values
func NewGetRecipientsOK() *GetRecipientsOK {

	return &GetRecipientsOK{}
}

// WithPayload adds the payload to the get recipients o k response
func (o *GetRecipientsOK) WithPayload(payload []*models.Recipient) *GetRecipientsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get recipients o k response
func (o *GetRecipientsOK) SetPayload(payload []*models.Recipient) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRecipientsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Recipient, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}