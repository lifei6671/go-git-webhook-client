package models

import "encoding/json"

type JsonResult struct {
	ErrorCode int                 	`json:"error_code"`
	Message string                	`json:"message"`
	Command string			`json:"command,omitempty"`
	MsgId string			`json:"msg_id,omitempty"`
	Data interface{}	      	`json:"data,omitempty"`
}

func (r *JsonResult) JsonString() (string,error) {
	b,err := json.Marshal(*r)
	if err != nil {
		return "",err
	}
	return string(b),nil
}

func FromString(s string) (*JsonResult,error) {
	var r JsonResult

	err := json.Unmarshal([]byte(s),&r)

	return &r,err
}
