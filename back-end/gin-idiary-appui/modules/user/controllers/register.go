package controllers

import (
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	userService "gin-idiary-appui/modules/user/services"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	inputs, hasError := register(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
	}
	err := userService.Register(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx)
}

func register(ctx *gin.Context) (userModel.RegisterPars, bool) {
	var inputs userModel.RegisterPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
