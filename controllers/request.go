package controllers

import "net/http"

type HttpRequest struct {
	request *http.Request
}

func (req *HttpRequest) ApplicationPath() string {
	return ""
}

func (req *HttpRequest) IsWebSocketRequest() bool {
	return true
}

func (req *HttpRequest) AcceptTypes() []string  {
	return make([]string,1)
}

func (req *HttpRequest) ContentEncoding() string  {
	return ""
}

func (req *HttpRequest) ContentLength() int64  {

	return 0;
}

func (req *HttpRequest) ContentType() string  {
	return ""
}

func (req *HttpRequest) Cookies()  {

}

func (req *HttpRequest) Files()  {

}

func (req *HttpRequest) Form(key string)  {

}

func (req *HttpRequest) Headers() {

}

func (req *HttpRequest) HttpMethod() {

}

func (req *HttpRequest) InputStream()  {

}

func (req *HttpRequest) Item(name string) {

}

func (req *HttpRequest) QueryString(name string) {

}

func (req *HttpRequest) Url() {

}

func (req *HttpRequest) UrlReferrer() {

}

func (req *HttpRequest) UserAgent() {

}

func (req *HttpRequest) UserHostAddress() {

}