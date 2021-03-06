// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2002Data inline response 200 2 data
// swagger:model inline_response_200_2_data
type InlineResponse2002Data struct {

	// members
	Members InlineResponse2002DataMembers `json:"members"`

	// 返回总数量
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this inline response 200 2 data
func (m *InlineResponse2002Data) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2002Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2002Data) UnmarshalBinary(b []byte) error {
	var res InlineResponse2002Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
