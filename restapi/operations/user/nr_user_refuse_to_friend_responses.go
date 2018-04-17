// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// UserRefuseToFriendOKCode is the HTTP code returned for type UserRefuseToFriendOK
const UserRefuseToFriendOKCode int = 200

/*UserRefuseToFriendOK 成功，返回成功信息

swagger:response userRefuseToFriendOK
*/
type UserRefuseToFriendOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2004 `json:"body,omitempty"`
}

// NewUserRefuseToFriendOK creates UserRefuseToFriendOK with default headers values
func NewUserRefuseToFriendOK() *UserRefuseToFriendOK {
	return &UserRefuseToFriendOK{}
}

// WithPayload adds the payload to the user refuse to friend o k response
func (o *UserRefuseToFriendOK) WithPayload(payload *models.InlineResponse2004) *UserRefuseToFriendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user refuse to friend o k response
func (o *UserRefuseToFriendOK) SetPayload(payload *models.InlineResponse2004) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserRefuseToFriendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrUserRefuseToFriendDefault 返回错误

swagger:response userRefuseToFriendDefault
*/
type NrUserRefuseToFriendDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrUserRefuseToFriendDefault creates NrUserRefuseToFriendDefault with default headers values
func NewNrUserRefuseToFriendDefault(code int) *NrUserRefuseToFriendDefault {
	if code <= 0 {
		code = 500
	}

	return &NrUserRefuseToFriendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user refuse to friend default response
func (o *NrUserRefuseToFriendDefault) WithStatusCode(code int) *NrUserRefuseToFriendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user refuse to friend default response
func (o *NrUserRefuseToFriendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user refuse to friend default response
func (o *NrUserRefuseToFriendDefault) WithPayload(payload *models.Error) *NrUserRefuseToFriendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user refuse to friend default response
func (o *NrUserRefuseToFriendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrUserRefuseToFriendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
