// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrUserShowParams creates a new NrUserShowParams object
// with the default values initialized.
func NewNrUserShowParams() NrUserShowParams {
	var ()
	return NrUserShowParams{}
}

// NrUserShowParams contains all the bound params for the user show operation
// typically these are obtained from a http.Request
//
// swagger:parameters /user/show
type NrUserShowParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*加密后的用户ID
	  Required: true
	  In: query
	*/
	Mid string
	/*
	  In: query
	*/
	Oid *string
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
func (o *NrUserShowParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qMid, qhkMid, _ := qs.GetOK("mid")
	if err := o.bindMid(qMid, qhkMid, route.Formats); err != nil {
		res = append(res, err)
	}

	qOid, qhkOid, _ := qs.GetOK("oid")
	if err := o.bindOid(qOid, qhkOid, route.Formats); err != nil {
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

func (o *NrUserShowParams) bindMid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrUserShowParams) bindOid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrUserShowParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrUserShowParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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