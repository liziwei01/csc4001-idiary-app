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
	}
	err := diaryService.AddDiary(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx)
}

func getAddDiaryPars(ctx *gin.Context) (diaryModel.DiaryRegisterPars, bool) {
	var inputs diaryModel.DiaryRegisterPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
