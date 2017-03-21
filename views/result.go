package views

import (
	"go-git-webhook-client/controllers"
)

type ViewResult interface {
	ExecuteResult (controllers.ControllerInterface)
}
