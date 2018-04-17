// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrUserDeleteFriendHandlerFunc turns a function with the right signature into a user delete friend handler
type NrUserDeleteFriendHandlerFunc func(NrUserDeleteFriendParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserDeleteFriendHandlerFunc) Handle(params NrUserDeleteFriendParams) middleware.Responder {
	return fn(params)
}

// NrUserDeleteFriendHandler interface for that can handle valid user delete friend params
type NrUserDeleteFriendHandler interface {
	Handle(NrUserDeleteFriendParams) middleware.Responder
}

// NewNrUserDeleteFriend creates a new http.Handler for the user delete friend operation
func NewNrUserDeleteFriend(ctx *middleware.Context, handler NrUserDeleteFriendHandler) *NrUserDeleteFriend {
	return &NrUserDeleteFriend{Context: ctx, Handler: handler}
}

/*NrUserDeleteFriend swagger:route POST /user/deleteFriend User userDeleteFriend

删除好友关系

*/
type NrUserDeleteFriend struct {
	Context *middleware.Context
	Handler NrUserDeleteFriendHandler
}

func (o *NrUserDeleteFriend) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserDeleteFriendParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
