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

// Body6 body 6
// swagger:model body_6
type Body6 struct {

	// 申请人的ID
	// Required: true
	Euid *string `json:"euid"`

	// oid
	Oid string `json:"oid,omitempty"`

	// 拒绝理由
	Reason string `json:"reason,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// 申请加入圈子的ID
	// Required: true
	ZoneID *int64 `json:"zone_id"`
}

// Validate validates this body 6
func (m *Body6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEuid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateZoneID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Body6) validateEuid(formats strfmt.Registry) error {

	if err := validate.Required("euid", "body", m.Euid); err != nil {
		return err
	}

	return nil
}

func (m *Body6) validateZoneID(formats strfmt.Registry) error {

	if err := validate.Required("zone_id", "body", m.ZoneID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Body6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body6) UnmarshalBinary(b []byte) error {
	var res Body6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}