/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:52:43
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:55:15
 * @Description: file content
 */
package comment

import (
	"gin-idiary-appui/library/response"
	commentModel "gin-idiary-appui/modules/diary/model/comment"
	commentService "gin-idiary-appui/modules/diary/services/comment"

	"github.com/gin-gonic/gin"
)

func Comment(ctx *gin.Context) {
	inputs, hasError := getCommentPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := commentService.Comment(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getCommentPars(ctx *gin.Context) (commentModel.CommentInfo, bool) {
	var inputs commentModel.CommentInfo
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
