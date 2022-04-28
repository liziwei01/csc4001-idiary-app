/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:52:43
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-24 20:37:22
 * @Description: file content
 */
package comment

import (
	"gin-idiary-appui/library/response"
	commentModel "gin-idiary-appui/modules/diary/model/comment"
	commentService "gin-idiary-appui/modules/diary/services/comment"

	"github.com/gin-gonic/gin"
)

func Like(ctx *gin.Context) {
	inputs, hasError := getLikePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := commentService.Like(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getLikePars(ctx *gin.Context) (commentModel.DiaryLikeRequestPars, bool) {
	var inputs commentModel.DiaryLikeRequestPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
