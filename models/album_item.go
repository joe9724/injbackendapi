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

// AlbumItem album item
// swagger:model AlbumItem
type AlbumItem struct {

	// 返回的日期(yyyy/MM/dd格式)
	// Required: true
	Date *string `json:"date"`

	// photos
	// Required: true
	Photos AlbumItemPhotos `json:"photos"`
}

// Validate validates this album item
func (m *AlbumItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhotos(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlbumItem) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

func (m *AlbumItem) validatePhotos(formats strfmt.Registry) error {

	if err := validate.Required("photos", "body", m.Photos); err != nil {
		return err
	}

	if err := m.Photos.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("photos")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlbumItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlbumItem) UnmarshalBinary(b []byte) error {
	var res AlbumItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
