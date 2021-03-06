// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrUserFansParams creates a new NrUserFansParams object
// with the default values initialized.
func NewNrUserFansParams() NrUserFansParams {
	var ()
	return NrUserFansParams{}
}

// NrUserFansParams contains all the bound params for the user fans operation
// typically these are obtained from a http.Request
//
// swagger:parameters /user/fans
type NrUserFansParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*单页返回的记录条数，默认为20
	  In: query
	*/
	Count *int64
	/*加密后的用户ID
	  Required: true
	  In: query
	*/
	Mid string
	/*
	  In: query
	*/
	Oid *string
	/*返回结果的页码，默认为1
	  Required: true
	  In: query
	*/
	Page int64
	/*
	  In: query
	*/
	Token *string
	/*加密后的用户ID
	  Required: true
	  In: query
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrUserFansParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCount, qhkCount, _ := qs.GetOK("count")
	if err := o.bindCount(qCount, qhkCount, route.Formats); err != nil {
		res = append(res, err)
	}

	qMid, qhkMid, _ := qs.GetOK("mid")
	if err := o.bindMid(qMid, qhkMid, route.Formats); err != nil {
		res = append(res, err)
	}

	qOid, qhkOid, _ := qs.GetOK("oid")
	if err := o.bindOid(qOid, qhkOid, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qToken, qhkToken, _ := qs.GetOK("token")
	if err := o.bindToken(qToken, qhkToken, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserID, qhkUserID, _ := qs.GetOK("user_id")
	if err := o.bindUserID(qUserID, qhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrUserFansParams) bindCount(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("count", "query", "int64", raw)
	}
	o.Count = &value

	return nil
}

func (o *NrUserFansParams) bindMid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("mid", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("mid", "query", raw); err != nil {
		return err
	}

	o.Mid = raw

	return nil
}

func (o *NrUserFansParams) bindOid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Oid = &raw

	return nil
}

func (o *NrUserFansParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("page", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("page", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = value

	return nil
}

func (o *NrUserFansParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Token = &raw

	return nil
}

func (o *NrUserFansParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("user_id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("user_id", "query", raw); err != nil {
		return err
	}

	o.UserID = raw

	return nil
}
