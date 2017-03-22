package controllers

import "time"

type ObjectCache interface {
	Insert(string,interface{},time.Duration)
}
