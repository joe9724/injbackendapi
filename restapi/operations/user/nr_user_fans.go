// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrUserFansHandlerFunc turns a function with the right signature into a user fans handler
type NrUserFansHandlerFunc func(NrUserFansParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserFansHandlerFunc) Handle(params NrUserFansParams) middleware.Responder {
	return fn(params)
}

// NrUserFansHandler interface for that can handle valid user fans params
type NrUserFansHandler interface {
	Handle(NrUserFansParams) middleware.Responder
}

// NewNrUserFans creates a new http.Handler for the user fans operation
func NewNrUserFans(ctx *middleware.Context, handler NrUserFansHandler) *NrUserFans {
	return &NrUserFans{Context: ctx, Handler: handler}
}

/*NrUserFans swagger:route GET /user/fans User userFans

查看用户的粉丝

*/
type NrUserFans struct {
	Context *middleware.Context
	Handler NrUserFansHandler
}

func (o *NrUserFans) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserFansParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
