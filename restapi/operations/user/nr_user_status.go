// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NrUserStatusHandlerFunc turns a function with the right signature into a user status handler
type NrUserStatusHandlerFunc func(NrUserStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserStatusHandlerFunc) Handle(params NrUserStatusParams) middleware.Responder {
	return fn(params)
}

// NrUserStatusHandler interface for that can handle valid user status params
type NrUserStatusHandler interface {
	Handle(NrUserStatusParams) middleware.Responder
}

// NewNrUserStatus creates a new http.Handler for the user status operation
func NewNrUserStatus(ctx *middleware.Context, handler NrUserStatusHandler) *NrUserStatus {
	return &NrUserStatus{Context: ctx, Handler: handler}
}

/*NrUserStatus swagger:route GET /user/status User userStatus

获取用户的动态

*/
type NrUserStatus struct {
	Context *middleware.Context
	Handler NrUserStatusHandler
}

func (o *NrUserStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok UserStatusOK
	var response models.InlineResponse20010
	var data models.InlineResponse20010Data
	var list []models.TimeLine
	//var temp []*models.ZoneItem

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	//var zoneList []models.ZoneItem
	db.Raw("select `when`,who,what from btk_Timeline").Find(&list)
	//fmt.Println("temp is",temp[0].Name)
	var count int64
	//db.Raw("select ZoneID from btk_Zone").Count(&count)
	//status
	data.TimeLines = list
	data.TotalCount = count
	response.Data = &data
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Response = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)
}
