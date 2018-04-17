// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrStatusCreateHandlerFunc turns a function with the right signature into a status create handler
type NrStatusCreateHandlerFunc func(NrStatusCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrStatusCreateHandlerFunc) Handle(params NrStatusCreateParams) middleware.Responder {
	return fn(params)
}

// NrStatusCreateHandler interface for that can handle valid status create params
type NrStatusCreateHandler interface {
	Handle(NrStatusCreateParams) middleware.Responder
}

// NewNrStatusCreate creates a new http.Handler for the status create operation
func NewNrStatusCreate(ctx *middleware.Context, handler NrStatusCreateHandler) *NrStatusCreate {
	return &NrStatusCreate{Context: ctx, Handler: handler}
}

/*NrStatusCreate swagger:route POST /status/create Status statusCreate

发布动态

*/
type NrStatusCreate struct {
	Context *middleware.Context
	Handler NrStatusCreateHandler
}

func (o *NrStatusCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrStatusCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}