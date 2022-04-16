/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:17:01
 * @Description: file content
 */
package follow

import (
	"gin-idiary-appui/library/response"
	userModel "gin-idiary-appui/modules/user/model"
	followService "gin-idiary-appui/modules/user/services/follow"

	"github.com/gin-gonic/gin"
)

func Followers(ctx *gin.Context) {
	inputs, hasError := getFollowersPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	followers, err := followService.Followers(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, gin.H{
		"followers": followers,
		"count":     len(followers),
	})
}

func getFollowersPars(ctx *gin.Context) (userModel.FollowerPars, bool) {
	var inputs userModel.FollowerPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
