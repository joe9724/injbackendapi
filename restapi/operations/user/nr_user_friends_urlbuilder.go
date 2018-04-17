// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// NrUserFriendsURL generates an URL for the user friends operation
type NrUserFriendsURL struct {
	Count  *int64
	Mid    string
	Oid    *string
	Page   int64
	Token  *string
	UserID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrUserFriendsURL) WithBasePath(bp string) *NrUserFriendsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrUserFriendsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *NrUserFriendsURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/user/friends"

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

	mid := o.Mid
	if mid != "" {
		qs.Set("mid", mid)
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

	var token string
	if o.Token != nil {
		token = *o.Token
	}
	if token != "" {
		qs.Set("token", token)
	}

	userID := o.UserID
	if userID != "" {
		qs.Set("user_id", userID)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *NrUserFriendsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *NrUserFriendsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *NrUserFriendsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on NrUserFriendsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on NrUserFriendsURL")
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
func (o *NrUserFriendsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
