package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ActionResult interface {
	ExecuteResult (http.ResponseWriter,*http.Request)
}

type JsonResult struct {
	data interface{}
}

func NewJsonResult(data interface{}) *JsonResult  {
	return &JsonResult{data:data}
}

func (self *JsonResult) ExecuteResult(w http.ResponseWriter,r *http.Request) {
	result,err := json.Marshal(self.data)

	if err != nil {
		fmt.Printf("%v",err)
		w.WriteHeader(500)
		fmt.Fprint(w,err.Error())
	}else {
		w.Header().Add("Content-Type","application/json")
		fmt.Fprint(w,string(result))
	}
}
