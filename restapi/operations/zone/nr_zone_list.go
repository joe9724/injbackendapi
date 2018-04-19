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

// NrZoneListHandlerFunc turns a function with the right signature into a zone list handler
type NrZoneListHandlerFunc func(NrZoneListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrZoneListHandlerFunc) Handle(params NrZoneListParams) middleware.Responder {
	return fn(params)
}

// NrZoneListHandler interface for that can handle valid zone list params
type NrZoneListHandler interface {
	Handle(NrZoneListParams) middleware.Responder
}

// NewNrZoneList creates a new http.Handler for the zone list operation
func NewNrZoneList(ctx *middleware.Context, handler NrZoneListHandler) *NrZoneList {
	return &NrZoneList{Context: ctx, Handler: handler}
}

/*NrZoneList swagger:route GET /zone/list Zone zoneList

获取圈子列表(按修改时间倒序排列)

*/
type NrZoneList struct {
	Context *middleware.Context
	Handler NrZoneListHandler
}

func (o *NrZoneList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrZoneListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok ZoneListOK
	var response models.InlineResponse200
	var data models.InlineResponse200Data
	var list models.InlineResponse200DataZones
	//var temp []*models.ZoneItem

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	//var zoneList []models.ZoneItem
	db.Raw("select ZoneID as zone_id,Level as level,Name as name,Status as status,Logo as logo,Photo as photo from btk_Zone order by ZoneID desc").Find(&list)
	//fmt.Println("temp is",temp[0].Name)
	var count int64
	db.Raw("select ZoneID from btk_Zone").Count(&count)
	//status
	data.Zones = list
	data.TotalCount = count
	response.Data = &data
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Response = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
