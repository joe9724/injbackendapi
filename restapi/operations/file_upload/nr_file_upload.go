// Code generated by go-swagger; DO NOT EDIT.

package file_upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"io/ioutil"
	"github.com/go-openapi/runtime/middleware"
	"fmt"
	_ "os"
	"runtime"
	"time"
	_"strings"
	"injbackendapi/models"
	"injbackendapi/var"
	"strconv"
)

// NrFileUploadHandlerFunc turns a function with the right signature into a file upload handler
type NrFileUploadHandlerFunc func(NrFileUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrFileUploadHandlerFunc) Handle(params NrFileUploadParams) middleware.Responder {
	return fn(params)
}

// NrFileUploadHandler interface for that can handle valid file upload params
type NrFileUploadHandler interface {
	Handle(NrFileUploadParams) middleware.Responder
}

// NewNrFileUpload creates a new http.Handler for the file upload operation
func NewNrFileUpload(ctx *middleware.Context, handler NrFileUploadHandler) *NrFileUpload {
	return &NrFileUpload{Context: ctx, Handler: handler}
}

/*NrFileUpload swagger:route POST /file/upload FileUpload fileUpload

文件上传

*/
type NrFileUpload struct {
	Context *middleware.Context
	Handler NrFileUploadHandler
}

func (o *NrFileUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrFileUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok FileUploadOK
	var response models.InlineResponse20014
	var status models.Response

	var msg string
	var code int64

	var filename string
	filename = strconv.FormatInt((time.Now().UnixNano()), 10)

	//fmt.Println("filename is", *Params.Filename)

	//fmt.Println("Param.type is", Params.Type)

	//var contentType string
	//var file []byte
	code = 200
	msg = "ok"

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	if (Params.Files != nil) {
		file, err := ioutil.ReadAll(Params.Files)
		if err != nil {
			fmt.Println("err upload:", err.Error())
		}
		fmt.Println("1")
		//save
		//var lower string
		//lower = strings.ToLower(contentType)
		//fmt.Println("lower is", lower)
		//fmt.Println("file.size is", len(file))
		if (true) {
			if (runtime.GOOS == "windows") {
				err1 := ioutil.WriteFile(filename+".jpg", file, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			} else {
				err1 := ioutil.WriteFile("/root/go/src/resource/image/icon/"+filename+".jpg", file, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			}
			response.Url = _var.GetResourceDomain("icon") + filename + ".jpg"
			fmt.Println("response.Url is",response.Url)
			code = 200
			msg = "ok"
		} else {
			fmt.Println("2")
			code = 401
			msg = "image format need jpg or png"
		}
	}else{
		fmt.Println("files is nil")
	}

	status.UnmarshalBinary([]byte(_var.Response200(code, msg)))
	response.Response = &status

	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)


}
