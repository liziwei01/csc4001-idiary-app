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
	}
	diary, err := diaryService.AllDiary(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx, diary)
}

func getAllDiaryPars(ctx *gin.Context) (diaryModel.DiaryShowPars, bool) {
	var inputs diaryModel.DiaryShowPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
