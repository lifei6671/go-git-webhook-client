package controllers

import "fmt"

type PayloadController struct {
	Controller
}

func (c *PayloadController) IndexAction()  {
	fmt.Fprint(c.Response,"aaaaaaa")
}