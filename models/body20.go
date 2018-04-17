// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Body20 body 20
// swagger:model body_20
type Body20 struct {

	// 关注人的ID
	// Required: true
	Mid *string `json:"mid"`

	// oid
	Oid string `json:"oid,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// 被关注人的ID
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this body 20
func (m *Body20) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Body20) validateMid(formats strfmt.Registry) error {

	if err := validate.Required("mid", "body", m.Mid); err != nil {
		return err
	}

	return nil
}

func (m *Body20) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Body20) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body20) UnmarshalBinary(b []byte) error {
	var res Body20
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
