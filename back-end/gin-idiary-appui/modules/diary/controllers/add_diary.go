/*
 * @Author: liziwei01
 * @Date: 2022-04-09 23:52:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:35:29
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	diaryModel "gin-idiary-appui/modules/diary/model"
	diaryService "gin-idiary-appui/modules/diary/services"

	"github.com/gin-gonic/gin"
)

func AddDiary(ctx *gin.Context) {
	inputs, hasError := getAddDiaryPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := diaryService.AddDiary(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getAddDiaryPars(ctx *gin.Context) (diaryModel.DiaryInfo, bool) {
	var inputs diaryModel.DiaryInfo
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
