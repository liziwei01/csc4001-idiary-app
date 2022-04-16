/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 18:22:56
 * @Description: file content
 */
package follow

import (
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	followService "gin-idiary-appui/modules/user/services/follow"

	"github.com/gin-gonic/gin"
)

func Follow(ctx *gin.Context) {
	inputs, hasError := getAddFriendPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := followService.Follow(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getAddFriendPars(ctx *gin.Context) (userModel.FollowPars, bool) {
	var inputs userModel.FollowPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
