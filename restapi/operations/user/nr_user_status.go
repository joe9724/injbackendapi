// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrUserStatusHandlerFunc turns a function with the right signature into a user status handler
type NrUserStatusHandlerFunc func(NrUserStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserStatusHandlerFunc) Handle(params NrUserStatusParams) middleware.Responder {
	return fn(params)
}

// NrUserStatusHandler interface for that can handle valid user status params
type NrUserStatusHandler interface {
	Handle(NrUserStatusParams) middleware.Responder
}

// NewNrUserStatus creates a new http.Handler for the user status operation
func NewNrUserStatus(ctx *middleware.Context, handler NrUserStatusHandler) *NrUserStatus {
	return &NrUserStatus{Context: ctx, Handler: handler}
}

/*NrUserStatus swagger:route GET /user/status User userStatus

获取用户的动态

*/
type NrUserStatus struct {
	Context *middleware.Context
	Handler NrUserStatusHandler
}

func (o *NrUserStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
