/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-14 00:11:18
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	userService "gin-idiary-appui/modules/user/services"

	"github.com/gin-gonic/gin"
)

func ModifyPassword(ctx *gin.Context) {
	inputs, hasError := getModifyPasswordPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := userService.ModifyPassword(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getModifyPasswordPars(ctx *gin.Context) (userModel.UserPars, bool) {
	var inputs userModel.UserPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}