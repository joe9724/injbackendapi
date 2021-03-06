// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// StatusCreateOKCode is the HTTP code returned for type StatusCreateOK
const StatusCreateOKCode int = 200

/*StatusCreateOK 成功，返回成功信息

swagger:response statusCreateOK
*/
type StatusCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2004 `json:"body,omitempty"`
}

// NewStatusCreateOK creates StatusCreateOK with default headers values
func NewStatusCreateOK() *StatusCreateOK {
	return &StatusCreateOK{}
}

// WithPayload adds the payload to the status create o k response
func (o *StatusCreateOK) WithPayload(payload *models.InlineResponse2004) *StatusCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the status create o k response
func (o *StatusCreateOK) SetPayload(payload *models.InlineResponse2004) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StatusCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrStatusCreateDefault 返回错误

swagger:response statusCreateDefault
*/
type NrStatusCreateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrStatusCreateDefault creates NrStatusCreateDefault with default headers values
func NewNrStatusCreateDefault(code int) *NrStatusCreateDefault {
	if code <= 0 {
		code = 500
	}

	return &NrStatusCreateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the status create default response
func (o *NrStatusCreateDefault) WithStatusCode(code int) *NrStatusCreateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the status create default response
func (o *NrStatusCreateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the status create default response
func (o *NrStatusCreateDefault) WithPayload(payload *models.Error) *NrStatusCreateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the status create default response
func (o *NrStatusCreateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrStatusCreateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
