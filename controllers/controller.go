package controllers

type ControllerInterface interface {
	Prepare()
}


type Controller struct {
	controllerName string
	actionName     string
}

func (c *Controller) Prepare() {

}