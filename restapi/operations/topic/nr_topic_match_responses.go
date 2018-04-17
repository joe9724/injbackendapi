// Code generated by go-swagger; DO NOT EDIT.

package topic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injbackendapi/models"
)

// TopicMatchOKCode is the HTTP code returned for type TopicMatchOK
const TopicMatchOKCode int = 200

/*TopicMatchOK 成功，返回成功信息

swagger:response topicMatchOK
*/
type TopicMatchOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2006 `json:"body,omitempty"`
}

// NewTopicMatchOK creates TopicMatchOK with default headers values
func NewTopicMatchOK() *TopicMatchOK {
	return &TopicMatchOK{}
}

// WithPayload adds the payload to the topic match o k response
func (o *TopicMatchOK) WithPayload(payload *models.InlineResponse2006) *TopicMatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the topic match o k response
func (o *TopicMatchOK) SetPayload(payload *models.InlineResponse2006) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TopicMatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrTopicMatchDefault 返回错误

swagger:response topicMatchDefault
*/
type NrTopicMatchDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrTopicMatchDefault creates NrTopicMatchDefault with default headers values
func NewNrTopicMatchDefault(code int) *NrTopicMatchDefault {
	if code <= 0 {
		code = 500
	}

	return &NrTopicMatchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the topic match default response
func (o *NrTopicMatchDefault) WithStatusCode(code int) *NrTopicMatchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the topic match default response
func (o *NrTopicMatchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the topic match default response
func (o *NrTopicMatchDefault) WithPayload(payload *models.Error) *NrTopicMatchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the topic match default response
func (o *NrTopicMatchDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrTopicMatchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}