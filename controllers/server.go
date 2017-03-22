package controllers

import "net/url"

type HttpServerUtility struct {

}

func (server *HttpServerUtility)UrlEncode(s string) string {
	return url.PathEscape(s)
}

func (server *HttpServerUtility) UrlDecode(s string) (string,error) {
	return url.PathUnescape(s)
}

func (server *HttpServerUtility)HtmlEncode(s string) string {
	return s
}

func (server *HttpServerUtility) HtmlDecode(s string) string  {
	return s
}

func (server *HttpServerUtility) MapPath(p string) string  {
	return p
}