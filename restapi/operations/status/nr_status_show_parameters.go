// Code generated by go-swagger; DO NOT EDIT.

package status

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

// NewNrStatusShowParams creates a new NrStatusShowParams object
// with the default values initialized.
func NewNrStatusShowParams() NrStatusShowParams {
	var ()
	return NrStatusShowParams{}
}

// NrStatusShowParams contains all the bound params for the status show operation
// typically these are obtained from a http.Request
//
// swagger:parameters /status/show
type NrStatusShowParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*加密后的用户ID
	  Required: true
	  In: query
	*/
	Euid string
	/*
	  In: query
	*/
	Oid *string
	/*动态ID
	  Required: true
	  In: query
	*/
	Pid int64
	/*
	  In: query
	*/
	Token *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrStatusShowParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEuid, qhkEuid, _ := qs.GetOK("euid")
	if err := o.bindEuid(qEuid, qhkEuid, route.Formats); err != nil {
		res = append(res, err)
	}

	qOid, qhkOid, _ := qs.GetOK("oid")
	if err := o.bindOid(qOid, qhkOid, route.Formats); err != nil {
		res = append(res, err)
	}

	qPid, qhkPid, _ := qs.GetOK("pid")
	if err := o.bindPid(qPid, qhkPid, route.Formats); err != nil {
		res = append(res, err)
	}

	qToken, qhkToken, _ := qs.GetOK("token")
	if err := o.bindToken(qToken, qhkToken, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrStatusShowParams) bindEuid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("euid", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("euid", "query", raw); err != nil {
		return err
	}

	o.Euid = raw

	return nil
}

func (o *NrStatusShowParams) bindOid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrStatusShowParams) bindPid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("pid", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("pid", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pid", "query", "int64", raw)
	}
	o.Pid = value

	return nil
}

func (o *NrStatusShowParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
