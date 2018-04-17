// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrZoneMembersHandlerFunc turns a function with the right signature into a zone members handler
type NrZoneMembersHandlerFunc func(NrZoneMembersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrZoneMembersHandlerFunc) Handle(params NrZoneMembersParams) middleware.Responder {
	return fn(params)
}

// NrZoneMembersHandler interface for that can handle valid zone members params
type NrZoneMembersHandler interface {
	Handle(NrZoneMembersParams) middleware.Responder
}

// NewNrZoneMembers creates a new http.Handler for the zone members operation
func NewNrZoneMembers(ctx *middleware.Context, handler NrZoneMembersHandler) *NrZoneMembers {
	return &NrZoneMembers{Context: ctx, Handler: handler}
}

/*NrZoneMembers swagger:route GET /zone/members Zone zoneMembers

获取圈子成员

获取圈子成员(0:普通成员 1:管理员 2:圈主)

*/
type NrZoneMembers struct {
	Context *middleware.Context
	Handler NrZoneMembersHandler
}

func (o *NrZoneMembers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrZoneMembersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}