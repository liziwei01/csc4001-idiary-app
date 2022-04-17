/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:13:52
 * @Description: file content
 */
package follow

import (
	"gin-idiary-appui/library/response"
	followModel "gin-idiary-appui/modules/user/model/follow"
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

func getFollowersPars(ctx *gin.Context) (followModel.FollowerPars, bool) {
	var inputs followModel.FollowerPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
