package controllers

import (
	"gin-idiary-appui/library/response"
	userService "gin-idiary-appui/modules/email/services"
	userModel "gin-idiary-appui/modules/user/model"

	"github.com/gin-gonic/gin"
)

func ModifyPassword(ctx *gin.Context) {
	inputs, hasError := getAddUserPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
	}
	err := userService.AddUser(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx)
}

func getAddUserPars(ctx *gin.Context) (userModel.UserRegisterPars, bool) {
	var inputs userModel.UserRegisterPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
