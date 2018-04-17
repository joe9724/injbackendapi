// Code generated by go-swagger; DO NOT EDIT.

package topic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// NrTopicStatusURL generates an URL for the topic status operation
type NrTopicStatusURL struct {
	Count *int64
	Euid  *string
	Oid   *string
	Page  int64
	Tid   int64
	Token *string
	Type  int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrTopicStatusURL) WithBasePath(bp string) *NrTopicStatusURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrTopicStatusURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *NrTopicStatusURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/topic/status"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/inj-backend/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var count string
	if o.Count != nil {
		count = swag.FormatInt64(*o.Count)
	}
	if count != "" {
		qs.Set("count", count)
	}

	var euid string
	if o.Euid != nil {
		euid = *o.Euid
	}
	if euid != "" {
		qs.Set("euid", euid)
	}

	var oid string
	if o.Oid != nil {
		oid = *o.Oid
	}
	if oid != "" {
		qs.Set("oid", oid)
	}

	page := swag.FormatInt64(o.Page)
	if page != "" {
		qs.Set("page", page)
	}

	tid := swag.FormatInt64(o.Tid)
	if tid != "" {
		qs.Set("tid", tid)
	}

	var token string
	if o.Token != nil {
		token = *o.Token
	}
	if token != "" {
		qs.Set("token", token)
	}

	typeVar := swag.FormatInt64(o.Type)
	if typeVar != "" {
		qs.Set("type", typeVar)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *NrTopicStatusURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *NrTopicStatusURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *NrTopicStatusURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on NrTopicStatusURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on NrTopicStatusURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *NrTopicStatusURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}