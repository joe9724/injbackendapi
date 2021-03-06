// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// ZoneRecommendOKCode is the HTTP code returned for type ZoneRecommendOK
const ZoneRecommendOKCode int = 200

/*ZoneRecommendOK 成功，返回成功信息

swagger:response zoneRecommendOK
*/
type ZoneRecommendOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse200 `json:"body,omitempty"`
}

// NewZoneRecommendOK creates ZoneRecommendOK with default headers values
func NewZoneRecommendOK() *ZoneRecommendOK {
	return &ZoneRecommendOK{}
}

// WithPayload adds the payload to the zone recommend o k response
func (o *ZoneRecommendOK) WithPayload(payload *models.InlineResponse200) *ZoneRecommendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone recommend o k response
func (o *ZoneRecommendOK) SetPayload(payload *models.InlineResponse200) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ZoneRecommendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrZoneRecommendDefault 返回错误

swagger:response zoneRecommendDefault
*/
type NrZoneRecommendDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrZoneRecommendDefault creates NrZoneRecommendDefault with default headers values
func NewNrZoneRecommendDefault(code int) *NrZoneRecommendDefault {
	if code <= 0 {
		code = 500
	}

	return &NrZoneRecommendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the zone recommend default response
func (o *NrZoneRecommendDefault) WithStatusCode(code int) *NrZoneRecommendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the zone recommend default response
func (o *NrZoneRecommendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the zone recommend default response
func (o *NrZoneRecommendDefault) WithPayload(payload *models.Error) *NrZoneRecommendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zone recommend default response
func (o *NrZoneRecommendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrZoneRecommendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
