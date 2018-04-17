// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "injbackendapi/models"
)

// NewNrUserDeleteFriendParams creates a new NrUserDeleteFriendParams object
// with the default values initialized.
func NewNrUserDeleteFriendParams() NrUserDeleteFriendParams {
	var ()
	return NrUserDeleteFriendParams{}
}

// NrUserDeleteFriendParams contains all the bound params for the user delete friend operation
// typically these are obtained from a http.Request
//
// swagger:parameters /user/deleteFriend
type NrUserDeleteFriendParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*以JSON的方式提交数据
	  In: body
	*/
	Body *models.Body21
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrUserDeleteFriendParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Body21
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
