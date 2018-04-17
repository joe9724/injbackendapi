// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrZoneReportHandlerFunc turns a function with the right signature into a zone report handler
type NrZoneReportHandlerFunc func(NrZoneReportParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrZoneReportHandlerFunc) Handle(params NrZoneReportParams) middleware.Responder {
	return fn(params)
}

// NrZoneReportHandler interface for that can handle valid zone report params
type NrZoneReportHandler interface {
	Handle(NrZoneReportParams) middleware.Responder
}

// NewNrZoneReport creates a new http.Handler for the zone report operation
func NewNrZoneReport(ctx *middleware.Context, handler NrZoneReportHandler) *NrZoneReport {
	return &NrZoneReport{Context: ctx, Handler: handler}
}

/*NrZoneReport swagger:route POST /zone/report Zone zoneReport

举报圈子

举报圈子操作

*/
type NrZoneReport struct {
	Context *middleware.Context
	Handler NrZoneReportHandler
}

func (o *NrZoneReport) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrZoneReportParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
