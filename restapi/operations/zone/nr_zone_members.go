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

	var ok ZoneMembersOK
	var response models.InlineResponse2002
	var data models.InlineResponse2002Data
	var list models.InlineResponse2002DataMembers
	var count int64
	//var temp []*models.ZoneItem

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	//var zoneList []models.ZoneItem
	db.Raw("select btk_User.nick_name as nickname,btk_User.avatar,btk_User.id as euid,btk_User.register_at as registerAt,btk_User.login_at as loginAt,gender as gender,level as level from btk_User").Find(&list)
	//fmt.Println("temp is",temp[0].Name)

	db.Raw("select ZoneID from btk_Zone").Count(&count)
	//status
	data.Members = list
	data.TotalCount = count
	response.Data = &data
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Response = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
