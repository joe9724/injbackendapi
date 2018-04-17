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

// Body14 body 14
// swagger:model body_14
type Body14 struct {

	// 用户ID
	// Required: true
	Euid *string `json:"euid"`

	// oid
	Oid string `json:"oid,omitempty"`

	// 动态ID
	StatusID int64 `json:"status_id,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this body 14
func (m *Body14) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEuid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Body14) validateEuid(formats strfmt.Registry) error {

	if err := validate.Required("euid", "body", m.Euid); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Body14) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body14) UnmarshalBinary(b []byte) error {
	var res Body14
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
