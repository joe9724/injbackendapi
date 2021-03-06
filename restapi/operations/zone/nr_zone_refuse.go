// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrZoneRefuseHandlerFunc turns a function with the right signature into a zone refuse handler
type NrZoneRefuseHandlerFunc func(NrZoneRefuseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrZoneRefuseHandlerFunc) Handle(params NrZoneRefuseParams) middleware.Responder {
	return fn(params)
}

// NrZoneRefuseHandler interface for that can handle valid zone refuse params
type NrZoneRefuseHandler interface {
	Handle(NrZoneRefuseParams) middleware.Responder
}

// NewNrZoneRefuse creates a new http.Handler for the zone refuse operation
func NewNrZoneRefuse(ctx *middleware.Context, handler NrZoneRefuseHandler) *NrZoneRefuse {
	return &NrZoneRefuse{Context: ctx, Handler: handler}
}

/*NrZoneRefuse swagger:route POST /zone/refuse Zone zoneRefuse

拒绝加入圈子

拒绝加入圈子操作(圈主或管理员)

*/
type NrZoneRefuse struct {
	Context *middleware.Context
	Handler NrZoneRefuseHandler
}

func (o *NrZoneRefuse) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrZoneRefuseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
