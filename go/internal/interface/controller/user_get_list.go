package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nakaaaa/go-clean-architecture/go/internal/interface/presenter"
	"github.com/nakaaaa/go-clean-architecture/go/internal/usecase"
)

func NewUserGetListController(uc usecase.UserGetListUseCase) Controller {
	return func(ctx *gin.Context) {
		input, err := ToUserGetListInputPort(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		output, err := uc.Execute(ctx.Request.Context(), input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		response, err := presenter.FromUserGetListOutputPort(output)
		if err != nil {
			logger.Errorf("fail to presenter.FromUserGetListOutputPort(): err=%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func ToUserGetListInputPort(ctx *gin.Context) (*usecase.UserGetListInputPort, error) {
	return &usecase.UserGetListInputPort{}, nil
}
