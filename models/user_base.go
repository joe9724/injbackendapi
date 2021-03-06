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

// UserBase 用户基本信息
// swagger:model UserBase
type UserBase struct {

	// 用户头像
	// Required: true
	Avatar *string `json:"avatar"`

	// 出生日期(yyyy-MM-dd HH:mm:ss)
	Birthday string `json:"birthday,omitempty"`

	// 加密后的用户ID
	// Required: true
	Euid *string `json:"euid"`

	// 用户性别(0:保密 1:男 2:女)
	Gender int64 `json:"gender,omitempty"`

	// 用户等级
	Level int64 `json:"level,omitempty"`

	// 用户昵称
	// Required: true
	Nickname *string `json:"nickname"`

	// 个人说明
	Remark string `json:"remark,omitempty"`
}

// Validate validates this user base
func (m *UserBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvatar(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEuid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNickname(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserBase) validateAvatar(formats strfmt.Registry) error {

	if err := validate.Required("avatar", "body", m.Avatar); err != nil {
		return err
	}

	return nil
}

func (m *UserBase) validateEuid(formats strfmt.Registry) error {

	if err := validate.Required("euid", "body", m.Euid); err != nil {
		return err
	}

	return nil
}

func (m *UserBase) validateNickname(formats strfmt.Registry) error {

	if err := validate.Required("nickname", "body", m.Nickname); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserBase) UnmarshalBinary(b []byte) error {
	var res UserBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
