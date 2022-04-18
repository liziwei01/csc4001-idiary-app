package controllers

import (
	"gin-idiary-appui/library/jwtoken"
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	userService "gin-idiary-appui/modules/user/services"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	inputs, hasError := getLoginPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	success, err := userService.Login(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}

	if !success {
		response.StdFailed(ctx, "用户名或密码错误")
		return
	}

	jwtToken, err := jwtoken.GenToken(inputs.Email)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, gin.H{
		"Authorization": jwtToken,
	})
}

func getLoginPars(ctx *gin.Context) (userModel.LoginPars, bool) {
	var inputs userModel.LoginPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
