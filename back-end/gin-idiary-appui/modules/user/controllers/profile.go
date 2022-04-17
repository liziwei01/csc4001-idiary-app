/*
 * @Author: liziwei01
 * @Date: 2022-04-17 15:52:49
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 15:57:14
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	userService "gin-idiary-appui/modules/user/services"

	"github.com/gin-gonic/gin"
)

func ModifyProfile(ctx *gin.Context) {
	inputs, hasError := getModifyProfilePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := userService.ModifyProfile(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getModifyProfilePars(ctx *gin.Context) (userModel.Profile, bool) {
	var inputs userModel.Profile
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
