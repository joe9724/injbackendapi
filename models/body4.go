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

// Body4 body 4
// swagger:model body_4
type Body4 struct {

	// 申请人的ID
	// Required: true
	Euid *string `json:"euid"`

	// oid
	Oid string `json:"oid,omitempty"`

	// 申请备注
	Remark string `json:"remark,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// 申请加入圈子的ID
	// Required: true
	ZoneID *int64 `json:"zone_id"`
}

// Validate validates this body 4
func (m *Body4) Validate(formats strfmt.Registry) error {
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

func (m *Body4) validateEuid(formats strfmt.Registry) error {

	if err := validate.Required("euid", "body", m.Euid); err != nil {
		return err
	}

	return nil
}

func (m *Body4) validateZoneID(formats strfmt.Registry) error {

	if err := validate.Required("zone_id", "body", m.ZoneID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Body4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body4) UnmarshalBinary(b []byte) error {
	var res Body4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
