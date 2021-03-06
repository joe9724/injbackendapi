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

// Body17 body 17
// swagger:model body_17
type Body17 struct {

	// 需要回复的评论ID
	// Required: true
	Cid *int64 `json:"cid"`

	// 回复评论内容，必须做URLencode，内容不超过140个汉字
	// Required: true
	Content *string `json:"content"`

	// 评论人的ID
	// Required: true
	Euid *string `json:"euid"`

	// oid
	Oid string `json:"oid,omitempty"`

	// 需要评论的主题ID
	// Required: true
	SourceID *int64 `json:"source_id"`

	// 需要评论主题类型
	// Required: true
	SourceType *int64 `json:"source_type"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this body 17
func (m *Body17) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEuid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Body17) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *Body17) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *Body17) validateEuid(formats strfmt.Registry) error {

	if err := validate.Required("euid", "body", m.Euid); err != nil {
		return err
	}

	return nil
}

func (m *Body17) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("source_id", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *Body17) validateSourceType(formats strfmt.Registry) error {

	if err := validate.Required("source_type", "body", m.SourceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Body17) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body17) UnmarshalBinary(b []byte) error {
	var res Body17
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
