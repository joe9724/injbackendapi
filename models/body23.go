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

// Body23 body 23
// swagger:model body_23
type Body23 struct {

	// 申请人的ID
	// Required: true
	Mid *string `json:"mid"`

	// oid
	Oid string `json:"oid,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// 被申请用户的ID
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this body 23
func (m *Body23) Validate(formats strfmt.Registry) error {
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

func (m *Body23) validateMid(formats strfmt.Registry) error {

	if err := validate.Required("mid", "body", m.Mid); err != nil {
		return err
	}

	return nil
}

func (m *Body23) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Body23) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body23) UnmarshalBinary(b []byte) error {
	var res Body23
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
