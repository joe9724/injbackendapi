// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrStatusSearchHandlerFunc turns a function with the right signature into a status search handler
type NrStatusSearchHandlerFunc func(NrStatusSearchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrStatusSearchHandlerFunc) Handle(params NrStatusSearchParams) middleware.Responder {
	return fn(params)
}

// NrStatusSearchHandler interface for that can handle valid status search params
type NrStatusSearchHandler interface {
	Handle(NrStatusSearchParams) middleware.Responder
}

// NewNrStatusSearch creates a new http.Handler for the status search operation
func NewNrStatusSearch(ctx *middleware.Context, handler NrStatusSearchHandler) *NrStatusSearch {
	return &NrStatusSearch{Context: ctx, Handler: handler}
}

/*NrStatusSearch swagger:route GET /status/search Status statusSearch

关键词搜索

*/
type NrStatusSearch struct {
	Context *middleware.Context
	Handler NrStatusSearchHandler
}

func (o *NrStatusSearch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrStatusSearchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
