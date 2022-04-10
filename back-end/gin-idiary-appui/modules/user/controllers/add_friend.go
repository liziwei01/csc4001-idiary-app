package controllers

import (
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	userService "gin-idiary-appui/modules/user/services"

	"github.com/gin-gonic/gin"
)

func AddFriend(ctx *gin.Context) {
	inputs, hasError := addFriend(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
	}
	err := userService.AddFriend(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx)
}

func addFriend(ctx *gin.Context) (userModel.FriendsPars, bool) {
	var inputs userModel.FriendsPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
