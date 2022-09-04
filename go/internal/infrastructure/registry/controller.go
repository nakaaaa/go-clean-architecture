package registry

import "github.com/nakaaaa/go-clean-architecture/go/internal/interface/controller"

type Controller struct {
	UserGetList controller.Controller
}

func NewController(
	userGetList controller.Controller,
) *Controller {
	return &Controller{
		UserGetList: userGetList,
	}
}
