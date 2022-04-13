/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-14 00:09:59
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	userService "gin-idiary-appui/modules/user/services"

	"github.com/gin-gonic/gin"
)

func AddFriend(ctx *gin.Context) {
	inputs, hasError := getAddFriendPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := userService.AddFriend(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getAddFriendPars(ctx *gin.Context) (userModel.FriendsPars, bool) {
	var inputs userModel.FriendsPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
