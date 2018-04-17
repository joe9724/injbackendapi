// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2008 inline response 200 8
// swagger:model inline_response_200_8
type InlineResponse2008 struct {

	// post
	Post *StatusItem `json:"post,omitempty"`

	// response
	Response *Response `json:"response,omitempty"`
}

// Validate validates this inline response 200 8
func (m *InlineResponse2008) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2008) validatePost(formats strfmt.Registry) error {

	if swag.IsZero(m.Post) { // not required
		return nil
	}

	if m.Post != nil {

		if err := m.Post.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2008) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {

		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2008) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2008) UnmarshalBinary(b []byte) error {
	var res InlineResponse2008
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}