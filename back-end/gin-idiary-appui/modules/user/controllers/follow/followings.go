/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:16:09
 * @Description: file content
 */
package follow

import (
	"gin-idiary-appui/library/response"
	followModel "gin-idiary-appui/modules/user/model/follow"
	followService "gin-idiary-appui/modules/user/services/follow"

	"github.com/gin-gonic/gin"
)

func Followings(ctx *gin.Context) {
	inputs, hasError := getFollowingsPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	followings, err := followService.Followings(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, gin.H{
		"followings": followings,
		"count":      len(followings),
	})
}

func getFollowingsPars(ctx *gin.Context) (followModel.FollowingPars, bool) {
	var inputs followModel.FollowingPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
