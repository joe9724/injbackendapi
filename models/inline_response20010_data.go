// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20010Data inline response 200 10 data
// swagger:model inline_response_200_10_data
type InlineResponse20010Data struct {

	// status
	Status InlineResponse20010DataStatus `json:"status"`

	// 返回总数量
	TotalCount int64 `json:"total_count,omitempty"`

	TimeLines []TimeLine `json:"timelines"`
}

// Validate validates this inline response 200 10 data
func (m *InlineResponse20010Data) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20010Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20010Data) UnmarshalBinary(b []byte) error {
	var res InlineResponse20010Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
