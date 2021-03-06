// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrUserFollowHandlerFunc turns a function with the right signature into a user follow handler
type NrUserFollowHandlerFunc func(NrUserFollowParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserFollowHandlerFunc) Handle(params NrUserFollowParams) middleware.Responder {
	return fn(params)
}

// NrUserFollowHandler interface for that can handle valid user follow params
type NrUserFollowHandler interface {
	Handle(NrUserFollowParams) middleware.Responder
}

// NewNrUserFollow creates a new http.Handler for the user follow operation
func NewNrUserFollow(ctx *middleware.Context, handler NrUserFollowHandler) *NrUserFollow {
	return &NrUserFollow{Context: ctx, Handler: handler}
}

/*NrUserFollow swagger:route POST /user/follow User userFollow

关注用户

*/
type NrUserFollow struct {
	Context *middleware.Context
	Handler NrUserFollowHandler
}

func (o *NrUserFollow) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserFollowParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
