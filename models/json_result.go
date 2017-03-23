package models

import "encoding/json"

type JsonResult struct {
	ErrorCode int                 `json:"error_code"`
	Message string                `json:"message"`
	Data interface{}	      `json:"data,omitempty"`
}

func (r *JsonResult) JsonString() (string,error) {
	b,err := json.Marshal(*r)
	if err != nil {
		return "",err
	}
	return string(b),nil
}

