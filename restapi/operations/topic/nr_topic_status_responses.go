// Code generated by go-swagger; DO NOT EDIT.

package topic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// TopicStatusOKCode is the HTTP code returned for type TopicStatusOK
const TopicStatusOKCode int = 200

/*TopicStatusOK 成功，返回成功信息

swagger:response topicStatusOK
*/
type TopicStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2003 `json:"body,omitempty"`
}

// NewTopicStatusOK creates TopicStatusOK with default headers values
func NewTopicStatusOK() *TopicStatusOK {
	return &TopicStatusOK{}
}

// WithPayload adds the payload to the topic status o k response
func (o *TopicStatusOK) WithPayload(payload *models.InlineResponse2003) *TopicStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the topic status o k response
func (o *TopicStatusOK) SetPayload(payload *models.InlineResponse2003) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TopicStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrTopicStatusDefault 返回错误

swagger:response topicStatusDefault
*/
type NrTopicStatusDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrTopicStatusDefault creates NrTopicStatusDefault with default headers values
func NewNrTopicStatusDefault(code int) *NrTopicStatusDefault {
	if code <= 0 {
		code = 500
	}

	return &NrTopicStatusDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the topic status default response
func (o *NrTopicStatusDefault) WithStatusCode(code int) *NrTopicStatusDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the topic status default response
func (o *NrTopicStatusDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the topic status default response
func (o *NrTopicStatusDefault) WithPayload(payload *models.Error) *NrTopicStatusDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the topic status default response
func (o *NrTopicStatusDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrTopicStatusDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
