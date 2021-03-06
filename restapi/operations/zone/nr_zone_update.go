// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrZoneUpdateHandlerFunc turns a function with the right signature into a zone update handler
type NrZoneUpdateHandlerFunc func(NrZoneUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrZoneUpdateHandlerFunc) Handle(params NrZoneUpdateParams) middleware.Responder {
	return fn(params)
}

// NrZoneUpdateHandler interface for that can handle valid zone update params
type NrZoneUpdateHandler interface {
	Handle(NrZoneUpdateParams) middleware.Responder
}

// NewNrZoneUpdate creates a new http.Handler for the zone update operation
func NewNrZoneUpdate(ctx *middleware.Context, handler NrZoneUpdateHandler) *NrZoneUpdate {
	return &NrZoneUpdate{Context: ctx, Handler: handler}
}

/*NrZoneUpdate swagger:route POST /zone/update Zone zoneUpdate

修改圈子信息

修改圈子信息(圈主/管理员)

*/
type NrZoneUpdate struct {
	Context *middleware.Context
	Handler NrZoneUpdateHandler
}

func (o *NrZoneUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrZoneUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
