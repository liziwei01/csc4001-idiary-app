package controllers

import (
	"csc4001-idiary-app/back-end/gin-idiary-appui/library/response"
	diaryModel "csc4001-idiary-app/back-end/gin-idiary-appui/modules/diary/model"
	diaryService "csc4001-idiary-app/back-end/gin-idiary-appui/modules/diary/services"

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
