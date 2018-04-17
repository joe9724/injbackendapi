// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// ZoneListOKCode is the HTTP code returned for type ZoneListOK
const ZoneListOKCode int = 200

/*ZoneListOK 成功，返回成功信息

swagger:response zoneListOK
*/
type ZoneListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse200 `json:"body,omitempty"`
}

// NewZoneListOK creates ZoneListOK with default headers values
func NewZoneListOK() *ZoneListOK {
	return &ZoneListOK{}
}

// WithPayload adds the payload to the zone list o k response
func (o *ZoneListOK) WithPayload(payload *models.InlineResponse200) *ZoneListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone list o k response
func (o *ZoneListOK) SetPayload(payload *models.InlineResponse200) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ZoneListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrZoneListDefault 返回错误

swagger:response zoneListDefault
*/
type NrZoneListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrZoneListDefault creates NrZoneListDefault with default headers values
func NewNrZoneListDefault(code int) *NrZoneListDefault {
	if code <= 0 {
		code = 500
	}

	return &NrZoneListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the zone list default response
func (o *NrZoneListDefault) WithStatusCode(code int) *NrZoneListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the zone list default response
func (o *NrZoneListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the zone list default response
func (o *NrZoneListDefault) WithPayload(payload *models.Error) *NrZoneListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone list default response
func (o *NrZoneListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrZoneListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}