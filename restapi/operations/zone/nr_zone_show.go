// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	middleware "github.com/go-openapi/runtime/middleware"
	"injbackendapi/models"
	"fmt"
	"injbackendapi/var"
)

// NrZoneShowHandlerFunc turns a function with the right signature into a zone show handler
type NrZoneShowHandlerFunc func(NrZoneShowParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrZoneShowHandlerFunc) Handle(params NrZoneShowParams) middleware.Responder {
	return fn(params)
}

// NrZoneShowHandler interface for that can handle valid zone show params
type NrZoneShowHandler interface {
	Handle(NrZoneShowParams) middleware.Responder
}

// NewNrZoneShow creates a new http.Handler for the zone show operation
func NewNrZoneShow(ctx *middleware.Context, handler NrZoneShowHandler) *NrZoneShow {
	return &NrZoneShow{Context: ctx, Handler: handler}
}

/*NrZoneShow swagger:route GET /zone/show Zone zoneShow

获取圈子详情

获取圈子详情(包含圈子成员个数、标签、最新公告)

*/
type NrZoneShow struct {
	Context *middleware.Context
	Handler NrZoneShowHandler
}

func (o *NrZoneShow) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrZoneShowParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok ZoneShowOK
	var response models.InlineResponse2001
	var data models.ZomeInfo
	//var list models.InlineResponse200DataZones
	//var temp []*models.ZoneItem

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	//var zoneList []models.ZoneItem
	db.Raw("select Name as name,Brief as brief,Logo as logo,Photo as photo,IsPublic as isPublic,Status as status,Level as level from btk_Zone").First(&data)
	//fmt.Println("temp is",temp[0].Name)
	//var count int64
	//db.Raw("select ZoneID from btk_Zone").Count(&count)
	//status
	//data. = list
	//data.TotalCount = count
	response.Data = &data
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Response = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
