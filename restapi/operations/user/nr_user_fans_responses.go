// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// UserFansOKCode is the HTTP code returned for type UserFansOK
const UserFansOKCode int = 200

/*UserFansOK 成功，返回成功信息

swagger:response userFansOK
*/
type UserFansOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20013 `json:"body,omitempty"`
}

// NewUserFansOK creates UserFansOK with default headers values
func NewUserFansOK() *UserFansOK {
	return &UserFansOK{}
}

// WithPayload adds the payload to the user fans o k response
func (o *UserFansOK) WithPayload(payload *models.InlineResponse20013) *UserFansOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user fans o k response
func (o *UserFansOK) SetPayload(payload *models.InlineResponse20013) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserFansOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrUserFansDefault 返回错误

swagger:response userFansDefault
*/
type NrUserFansDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrUserFansDefault creates NrUserFansDefault with default headers values
func NewNrUserFansDefault(code int) *NrUserFansDefault {
	if code <= 0 {
		code = 500
	}

	return &NrUserFansDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user fans default response
func (o *NrUserFansDefault) WithStatusCode(code int) *NrUserFansDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user fans default response
func (o *NrUserFansDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user fans default response
func (o *NrUserFansDefault) WithPayload(payload *models.Error) *NrUserFansDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user fans default response
func (o *NrUserFansDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrUserFansDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}