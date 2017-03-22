package controllers

import (
	"net/http"
	"bytes"
	"fmt"
)

type HttpResponse struct {
	response http.ResponseWriter
	buffer *bytes.Buffer
	Status int
}

func (res *HttpResponse) Cookies() {

}

func (res *HttpResponse) Headers() {

}

func (res *HttpResponse) AddHeader(name string,value string) {

}

func (r *HttpResponse) ClearContent() {

}

func (r *HttpResponse) ClearHeaders() {

}

func (r *HttpResponse) Redirect(uri string) {

}

func (r *HttpResponse) WriteBytes(b []byte) {
	r.buffer.Write(b)
}
func (r *HttpResponse) WriteString(s string) {
	r.buffer.WriteString(s)
}

func (r *HttpResponse) End() {
	fmt.Fprint(r.response,r.buffer)
	panic("Response end.")
}