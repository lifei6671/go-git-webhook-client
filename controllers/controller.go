package controllers

import "net/http"


type ControllerInterface interface {
	Init(http.ResponseWriter,*http.Request)
}


type Controller struct {
	Request *http.Request
	Response http.ResponseWriter
}

func (c *Controller) Init(w http.ResponseWriter,r *http.Request) {
	c.Request = r
	c.Response = w
}