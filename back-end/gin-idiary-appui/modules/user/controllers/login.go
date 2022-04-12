package controllers

import (
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	userService "gin-idiary-appui/modules/user/services"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	inputs, hasError := getLoginPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
	}
	bool, err := userService.Login(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx, bool)
}

func getLoginPars(ctx *gin.Context) (userModel.LoginPars, bool) {
	var inputs userModel.LoginPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
