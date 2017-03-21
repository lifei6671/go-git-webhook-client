package controllers

import "fmt"

type PayloadController struct {
	Controller
}

func (c *PayloadController) IndexAction() {
	fmt.Println("aaaaaaaaaaaaaa")
}