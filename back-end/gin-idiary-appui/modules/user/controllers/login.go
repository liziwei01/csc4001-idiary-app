/*
 * @Author: liziwei01
 * @Date: 2022-04-18 17:27:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-18 20:51:42
 * @Description: file content
 */
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
	info, err := userService.Login(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}

	if info.Password != inputs.Password {
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
		"email":         info.Email,
		"user_id":       info.UserID,
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
