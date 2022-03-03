/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:18:46
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-03 20:09:57
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"

	userModel "gin-idiary-appui/modules/user/model"
	userService "gin-idiary-appui/modules/user/services"

	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context) {
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
