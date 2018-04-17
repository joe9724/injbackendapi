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

// Body3 body 3
// swagger:model body_3
type Body3 struct {

	// 邀请人的ID
	// Required: true
	Euid *string `json:"euid"`

	// oid
	Oid string `json:"oid,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// 被邀请的用户ID，可以传入多个用户id, 以","分隔
	// Required: true
	UserIds *string `json:"user_ids"`

	// 加入圈子的ID
	// Required: true
	ZoneID *int64 `json:"zone_id"`
}

// Validate validates this body 3
func (m *Body3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEuid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserIds(formats); err != nil {
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

func (m *Body3) validateEuid(formats strfmt.Registry) error {

	if err := validate.Required("euid", "body", m.Euid); err != nil {
		return err
	}

	return nil
}

func (m *Body3) validateUserIds(formats strfmt.Registry) error {

	if err := validate.Required("user_ids", "body", m.UserIds); err != nil {
		return err
	}

	return nil
}

func (m *Body3) validateZoneID(formats strfmt.Registry) error {

	if err := validate.Required("zone_id", "body", m.ZoneID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Body3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body3) UnmarshalBinary(b []byte) error {
	var res Body3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}