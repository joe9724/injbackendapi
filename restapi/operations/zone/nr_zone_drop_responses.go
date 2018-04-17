// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// ZoneDropOKCode is the HTTP code returned for type ZoneDropOK
const ZoneDropOKCode int = 200

/*ZoneDropOK 成功，返回成功信息

swagger:response zoneDropOK
*/
type ZoneDropOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2004 `json:"body,omitempty"`
}

// NewZoneDropOK creates ZoneDropOK with default headers values
func NewZoneDropOK() *ZoneDropOK {
	return &ZoneDropOK{}
}

// WithPayload adds the payload to the zone drop o k response
func (o *ZoneDropOK) WithPayload(payload *models.InlineResponse2004) *ZoneDropOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone drop o k response
func (o *ZoneDropOK) SetPayload(payload *models.InlineResponse2004) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ZoneDropOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrZoneDropDefault 返回错误

swagger:response zoneDropDefault
*/
type NrZoneDropDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrZoneDropDefault creates NrZoneDropDefault with default headers values
func NewNrZoneDropDefault(code int) *NrZoneDropDefault {
	if code <= 0 {
		code = 500
	}

	return &NrZoneDropDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the zone drop default response
func (o *NrZoneDropDefault) WithStatusCode(code int) *NrZoneDropDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the zone drop default response
func (o *NrZoneDropDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the zone drop default response
func (o *NrZoneDropDefault) WithPayload(payload *models.Error) *NrZoneDropDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone drop default response
func (o *NrZoneDropDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrZoneDropDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}