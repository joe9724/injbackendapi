// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// UserFollowsOKCode is the HTTP code returned for type UserFollowsOK
const UserFollowsOKCode int = 200

/*UserFollowsOK 成功，返回成功信息

swagger:response userFollowsOK
*/
type UserFollowsOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20013 `json:"body,omitempty"`
}

// NewUserFollowsOK creates UserFollowsOK with default headers values
func NewUserFollowsOK() *UserFollowsOK {
	return &UserFollowsOK{}
}

// WithPayload adds the payload to the user follows o k response
func (o *UserFollowsOK) WithPayload(payload *models.InlineResponse20013) *UserFollowsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user follows o k response
func (o *UserFollowsOK) SetPayload(payload *models.InlineResponse20013) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserFollowsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrUserFollowsDefault 返回错误

swagger:response userFollowsDefault
*/
type NrUserFollowsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrUserFollowsDefault creates NrUserFollowsDefault with default headers values
func NewNrUserFollowsDefault(code int) *NrUserFollowsDefault {
	if code <= 0 {
		code = 500
	}

	return &NrUserFollowsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user follows default response
func (o *NrUserFollowsDefault) WithStatusCode(code int) *NrUserFollowsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user follows default response
func (o *NrUserFollowsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user follows default response
func (o *NrUserFollowsDefault) WithPayload(payload *models.Error) *NrUserFollowsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user follows default response
func (o *NrUserFollowsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrUserFollowsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
