// Code generated by go-swagger; DO NOT EDIT.

package comment

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

// NewNrCommentsListParams creates a new NrCommentsListParams object
// with the default values initialized.
func NewNrCommentsListParams() NrCommentsListParams {
	var ()
	return NrCommentsListParams{}
}

// NrCommentsListParams contains all the bound params for the comments list operation
// typically these are obtained from a http.Request
//
// swagger:parameters /comments/list
type NrCommentsListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*单页返回的记录条数，默认为20
	  In: query
	*/
	Count *int64
	/*
	  In: query
	*/
	Keyword *string
	/*
	  In: query
	*/
	Oid *string
	/*返回结果的页码，默认为1
	  Required: true
	  In: query
	*/
	Page int64
	/*主题ID
	  Required: true
	  In: query
	*/
	SourceID int64
	/*主题类型(0:活动 1:动态 2:商品)
	  Required: true
	  In: query
	*/
	SourceType int64
	/*
	  In: query
	*/
	Token *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrCommentsListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCount, qhkCount, _ := qs.GetOK("count")
	if err := o.bindCount(qCount, qhkCount, route.Formats); err != nil {
		res = append(res, err)
	}

	qKeyword, qhkKeyword, _ := qs.GetOK("keyword")
	if err := o.bindKeyword(qKeyword, qhkKeyword, route.Formats); err != nil {
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

	qSourceID, qhkSourceID, _ := qs.GetOK("source_id")
	if err := o.bindSourceID(qSourceID, qhkSourceID, route.Formats); err != nil {
		res = append(res, err)
	}

	qSourceType, qhkSourceType, _ := qs.GetOK("source_type")
	if err := o.bindSourceType(qSourceType, qhkSourceType, route.Formats); err != nil {
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

func (o *NrCommentsListParams) bindCount(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrCommentsListParams) bindKeyword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Keyword = &raw

	return nil
}

func (o *NrCommentsListParams) bindOid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrCommentsListParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrCommentsListParams) bindSourceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("source_id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("source_id", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("source_id", "query", "int64", raw)
	}
	o.SourceID = value

	return nil
}

func (o *NrCommentsListParams) bindSourceType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("source_type", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("source_type", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("source_type", "query", "int64", raw)
	}
	o.SourceType = value

	return nil
}

func (o *NrCommentsListParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
