// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2003Data inline response 200 3 data
// swagger:model inline_response_200_3_data
type InlineResponse2003Data struct {

	// posts
	Posts InlineResponse2003DataPosts `json:"posts"`

	// 返回总数量
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this inline response 200 3 data
func (m *InlineResponse2003Data) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2003Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2003Data) UnmarshalBinary(b []byte) error {
	var res InlineResponse2003Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
