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

// StatusItem status item
// swagger:model StatusItem
type StatusItem struct {

	// ait users
	AitUsers StatusItemAitUsers `json:"ait_users"`

	// belongs forum
	BelongsForum *ZoneItem `json:"belongs_forum,omitempty"`

	// 评论数
	CommentsCount int64 `json:"comments_count,omitempty"`

	// 动态创建日期
	// Required: true
	CreatedAt *int64 `json:"created_at"`

	// geo
	Geo *GeoInfo `json:"geo,omitempty"`

	// 图片地址数组, 列表中最大6张
	Pics []string `json:"pics"`

	// 动态ID
	// Required: true
	Pid *int64 `json:"pid"`

	// 转发数
	RepostsCount int64 `json:"reposts_count,omitempty"`

	// retweeted status
	RetweetedStatus *RetweetedStatus `json:"retweeted_status,omitempty"`

	// 点赞数
	SupportCount int64 `json:"support_count,omitempty"`

	// 是否点赞
	Supported bool `json:"supported,omitempty"`

	// 动态内容(包含@的人、话题名称等)
	Text string `json:"text,omitempty"`

	// user
	// Required: true
	User *UserItem `json:"user"`
}

// Validate validates this status item
func (m *StatusItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBelongsForum(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGeo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePics(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetweetedStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusItem) validateBelongsForum(formats strfmt.Registry) error {

	if swag.IsZero(m.BelongsForum) { // not required
		return nil
	}

	if m.BelongsForum != nil {

		if err := m.BelongsForum.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("belongs_forum")
			}
			return err
		}
	}

	return nil
}

func (m *StatusItem) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *StatusItem) validateGeo(formats strfmt.Registry) error {

	if swag.IsZero(m.Geo) { // not required
		return nil
	}

	if m.Geo != nil {

		if err := m.Geo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geo")
			}
			return err
		}
	}

	return nil
}

func (m *StatusItem) validatePics(formats strfmt.Registry) error {

	if swag.IsZero(m.Pics) { // not required
		return nil
	}

	return nil
}

func (m *StatusItem) validatePid(formats strfmt.Registry) error {

	if err := validate.Required("pid", "body", m.Pid); err != nil {
		return err
	}

	return nil
}

func (m *StatusItem) validateRetweetedStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.RetweetedStatus) { // not required
		return nil
	}

	if m.RetweetedStatus != nil {

		if err := m.RetweetedStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retweeted_status")
			}
			return err
		}
	}

	return nil
}

func (m *StatusItem) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if m.User != nil {

		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusItem) UnmarshalBinary(b []byte) error {
	var res StatusItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}