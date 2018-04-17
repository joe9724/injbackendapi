// Code generated by go-swagger; DO NOT EDIT.

package file_upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrFileUploadParams creates a new NrFileUploadParams object
// with the default values initialized.
func NewNrFileUploadParams() NrFileUploadParams {
	var ()
	return NrFileUploadParams{}
}

// NrFileUploadParams contains all the bound params for the file upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters /file/upload
type NrFileUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*上传者的ID
	  In: formData
	*/
	Euid *string
	/*上传的文件数组
	  Required: true
	  In: formData
	*/
	Files io.ReadCloser
	/*
	  In: query
	*/
	Oid *string
	/*
	  In: query
	*/
	Token *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrFileUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	fdEuid, fdhkEuid, _ := fds.GetOK("euid")
	if err := o.bindEuid(fdEuid, fdhkEuid, route.Formats); err != nil {
		res = append(res, err)
	}

	files, filesHeader, err := r.FormFile("files")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "files", err))
	} else if err := o.bindFiles(files, filesHeader); err != nil {
		res = append(res, err)
	} else {
		o.Files = &runtime.File{Data: files, Header: filesHeader}
	}

	qOid, qhkOid, _ := qs.GetOK("oid")
	if err := o.bindOid(qOid, qhkOid, route.Formats); err != nil {
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

func (o *NrFileUploadParams) bindEuid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrFileUploadParams) bindFiles(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *NrFileUploadParams) bindOid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrFileUploadParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
