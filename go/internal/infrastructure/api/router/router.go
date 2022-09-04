package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/registry"
)

func Attach(group *gin.RouterGroup, controller *registry.Controller) {
	group.GET("healthcheck", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	AttachUsers(group, controller)
}

func AttachUsers(group *gin.RouterGroup, controller *registry.Controller) {
	group.GET("users", gin.HandlerFunc(controller.UserGetList))
}
