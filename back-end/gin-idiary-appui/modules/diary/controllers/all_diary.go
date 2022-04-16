/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-14 00:08:56
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	diaryModel "gin-idiary-appui/modules/diary/model"
	diaryService "gin-idiary-appui/modules/diary/services"

	"github.com/gin-gonic/gin"
)

func AllDiary(ctx *gin.Context) {
	inputs, hasError := getAllDiaryPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	diary, err := diaryService.AllDiary(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, diary)
}

func getAllDiaryPars(ctx *gin.Context) (diaryModel.DiaryListRequestPars, bool) {
	var inputs diaryModel.DiaryListRequestPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
