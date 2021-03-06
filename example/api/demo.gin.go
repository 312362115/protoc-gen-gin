// Code generated by protoc-gen-gin. DO NOT EDIT.
// source: api/demo.proto

package demo

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/312362115/protoc-gen-gin/ecode"
)

// to suppressed 'imported but not used warning'
var _ *gin.Context
var _ context.Context
var _ binding.StructValidator

var PathDemoHello = "/rpc/hello"
var PathDemoDemo = "/rpc/demo"

// DemoGinServer is the server API for Demo service.
type DemoGinServer interface {
	Hello(ctx context.Context, req *HelloReq) (resp *HelloResp, err error)

	Demo(ctx context.Context, req *DemoReq) (resp *DemoResp, err error)
}

func JSON(c *gin.Context, data interface{}, err error) {
	httpCode := http.StatusOK
	bcode := ecode.Cause(err)
	if bcode.Code < 0 {
		httpCode = -bcode.Code
	}
	c.JSON(httpCode, Response{
		Code:    bcode.Code,
		Message: bcode.Message,
		Data:    data,
	})
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var DemoSvc DemoGinServer

func demoHello(c *gin.Context) {
	p := new(HelloReq)

	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.Hello(c, p)
	JSON(c, resp, err)
}

func demoDemo(c *gin.Context) {
	p := new(DemoReq)

	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.Demo(c, p)
	JSON(c, resp, err)
}

// RegisterDemoGinServer Register the gin route
func RegisterDemoGinServer(c *gin.Engine, server DemoGinServer) {
	DemoSvc = server
	c.GET("/rpc/hello", demoHello)
	c.POST("/rpc/demo", demoDemo)
}
