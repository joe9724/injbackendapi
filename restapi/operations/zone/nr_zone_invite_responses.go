// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// ZoneInviteOKCode is the HTTP code returned for type ZoneInviteOK
const ZoneInviteOKCode int = 200

/*ZoneInviteOK 成功，返回成功信息

swagger:response zoneInviteOK
*/
type ZoneInviteOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2004 `json:"body,omitempty"`
}

// NewZoneInviteOK creates ZoneInviteOK with default headers values
func NewZoneInviteOK() *ZoneInviteOK {
	return &ZoneInviteOK{}
}

// WithPayload adds the payload to the zone invite o k response
func (o *ZoneInviteOK) WithPayload(payload *models.InlineResponse2004) *ZoneInviteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone invite o k response
func (o *ZoneInviteOK) SetPayload(payload *models.InlineResponse2004) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ZoneInviteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrZoneInviteDefault 返回错误

swagger:response zoneInviteDefault
*/
type NrZoneInviteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrZoneInviteDefault creates NrZoneInviteDefault with default headers values
func NewNrZoneInviteDefault(code int) *NrZoneInviteDefault {
	if code <= 0 {
		code = 500
	}

	return &NrZoneInviteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the zone invite default response
func (o *NrZoneInviteDefault) WithStatusCode(code int) *NrZoneInviteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the zone invite default response
func (o *NrZoneInviteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the zone invite default response
func (o *NrZoneInviteDefault) WithPayload(payload *models.Error) *NrZoneInviteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone invite default response
func (o *NrZoneInviteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrZoneInviteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}