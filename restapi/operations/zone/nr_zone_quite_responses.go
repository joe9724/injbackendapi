// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// ZoneQuiteOKCode is the HTTP code returned for type ZoneQuiteOK
const ZoneQuiteOKCode int = 200

/*ZoneQuiteOK 成功，返回成功信息

swagger:response zoneQuiteOK
*/
type ZoneQuiteOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2004 `json:"body,omitempty"`
}

// NewZoneQuiteOK creates ZoneQuiteOK with default headers values
func NewZoneQuiteOK() *ZoneQuiteOK {
	return &ZoneQuiteOK{}
}

// WithPayload adds the payload to the zone quite o k response
func (o *ZoneQuiteOK) WithPayload(payload *models.InlineResponse2004) *ZoneQuiteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone quite o k response
func (o *ZoneQuiteOK) SetPayload(payload *models.InlineResponse2004) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ZoneQuiteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrZoneQuiteDefault 返回错误

swagger:response zoneQuiteDefault
*/
type NrZoneQuiteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrZoneQuiteDefault creates NrZoneQuiteDefault with default headers values
func NewNrZoneQuiteDefault(code int) *NrZoneQuiteDefault {
	if code <= 0 {
		code = 500
	}

	return &NrZoneQuiteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the zone quite default response
func (o *NrZoneQuiteDefault) WithStatusCode(code int) *NrZoneQuiteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the zone quite default response
func (o *NrZoneQuiteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the zone quite default response
func (o *NrZoneQuiteDefault) WithPayload(payload *models.Error) *NrZoneQuiteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone quite default response
func (o *NrZoneQuiteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrZoneQuiteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
