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

// NewNrStatusTimelineParams creates a new NrStatusTimelineParams object
// with the default values initialized.
func NewNrStatusTimelineParams() NrStatusTimelineParams {
	var ()
	return NrStatusTimelineParams{}
}

// NrStatusTimelineParams contains all the bound params for the status timeline operation
// typically these are obtained from a http.Request
//
// swagger:parameters /status/timeline
type NrStatusTimelineParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*单页返回的记录条数，默认为20
	  In: query
	*/
	Count *int64
	/*加密后的用户ID
	  In: query
	*/
	Euid *string
	/*
	  In: query
	*/
	Keyword *string
	/*获取周边动态时需要
	  In: query
	*/
	Latitude *float32
	/*获取周边动态时需要
	  In: query
	*/
	Longitude *float32
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
	/*获取动态的类型(0:热门  1:关注  2:好友  3:推荐  4:周边)
	  Required: true
	  In: query
	*/
	Type int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrStatusTimelineParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCount, qhkCount, _ := qs.GetOK("count")
	if err := o.bindCount(qCount, qhkCount, route.Formats); err != nil {
		res = append(res, err)
	}

	qEuid, qhkEuid, _ := qs.GetOK("euid")
	if err := o.bindEuid(qEuid, qhkEuid, route.Formats); err != nil {
		res = append(res, err)
	}

	qKeyword, qhkKeyword, _ := qs.GetOK("keyword")
	if err := o.bindKeyword(qKeyword, qhkKeyword, route.Formats); err != nil {
		res = append(res, err)
	}

	qLatitude, qhkLatitude, _ := qs.GetOK("latitude")
	if err := o.bindLatitude(qLatitude, qhkLatitude, route.Formats); err != nil {
		res = append(res, err)
	}

	qLongitude, qhkLongitude, _ := qs.GetOK("longitude")
	if err := o.bindLongitude(qLongitude, qhkLongitude, route.Formats); err != nil {
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

	qType, qhkType, _ := qs.GetOK("type")
	if err := o.bindType(qType, qhkType, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrStatusTimelineParams) bindCount(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrStatusTimelineParams) bindEuid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Euid = &raw

	return nil
}

func (o *NrStatusTimelineParams) bindKeyword(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrStatusTimelineParams) bindLatitude(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat32(raw)
	if err != nil {
		return errors.InvalidType("latitude", "query", "float32", raw)
	}
	o.Latitude = &value

	return nil
}

func (o *NrStatusTimelineParams) bindLongitude(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat32(raw)
	if err != nil {
		return errors.InvalidType("longitude", "query", "float32", raw)
	}
	o.Longitude = &value

	return nil
}

func (o *NrStatusTimelineParams) bindOid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrStatusTimelineParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrStatusTimelineParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrStatusTimelineParams) bindType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("type", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("type", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("type", "query", "int64", raw)
	}
	o.Type = value

	return nil
}
