package controllers

import (
	"net/http"
	"bytes"
)


type ControllerInterface interface {
	Init(http.ResponseWriter,*http.Request)
}


type Controller struct {
	Request *HttpRequest
	Response *HttpResponse
	Server *HttpServerUtility
	Cache ObjectCache
	Session HttpSessionStateBase
	Items map[string]interface{}
}

func (c *Controller) Init(w http.ResponseWriter,r *http.Request) {
	c.Request = &HttpRequest{request : r}
	c.Response = &HttpResponse{ response : w, buffer : bytes.NewBufferString(""),Status : 200}
	c.Server = &HttpServerUtility{}
	c.Items = make(map[string]interface{},0)

}