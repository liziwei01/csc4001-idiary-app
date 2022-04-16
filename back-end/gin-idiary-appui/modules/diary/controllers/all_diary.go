/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:51:16
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
	diaries, count, err := diaryService.AllDiary(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, gin.H{
		"diaries": diaries,
		"count":   count,
	})
}

func getAllDiaryPars(ctx *gin.Context) (diaryModel.DiaryListRequestPars, bool) {
	var inputs diaryModel.DiaryListRequestPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	if inputs.PageIndex < 0 {
		inputs.PageIndex = 0
	}
	if inputs.PageLength <= 0 {
		inputs.PageLength = 10
	}
	return inputs, false
}
