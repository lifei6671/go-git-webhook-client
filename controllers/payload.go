package controllers

type PayloadController struct {
	Controller
}

func (c *PayloadController) IndexAction() ActionResult {

	c.Response.WriteString("aaa")


	return NewJsonResult(map[string]string{"Key":"aaaa"})
}