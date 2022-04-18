package controllers

import (
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	userService "gin-idiary-appui/modules/user/services"

	"github.com/gin-gonic/gin"
)

func AllUsers(ctx *gin.Context) {
	inputs, hasError := getAllUsersPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	users, count, err := userService.AllUsers(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, gin.H{
		"users": users,
		"count": count,
	})
}

func getAllUsersPars(ctx *gin.Context) (userModel.UserListRequestPars, bool) {
	var inputs userModel.UserListRequestPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	if inputs.PageIndex < 0 {
		inputs.PageIndex = 0
	}
	if inputs.PageLength <= 0 {
		inputs.PageLength = 10
	}
	return inputs, false
}
