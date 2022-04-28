/*
 * @Author: liziwei01
 * @Date: 2022-04-09 23:52:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-25 21:24:39
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	diaryModel "gin-idiary-appui/modules/diary/model"
	diaryService "gin-idiary-appui/modules/diary/services"

	"github.com/gin-gonic/gin"
)

func DeleteDiary(ctx *gin.Context) {
	inputs, hasError := getDeleteDiaryPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := diaryService.DeleteDiary(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getDeleteDiaryPars(ctx *gin.Context) (diaryModel.DeleteDiaryRequestPars, bool) {
	var inputs diaryModel.DeleteDiaryRequestPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
