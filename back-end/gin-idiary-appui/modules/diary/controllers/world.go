package controllers

import (
	"gin-idiary-appui/library/response"
	diaryModel "gin-idiary-appui/modules/diary/model"
	diaryService "gin-idiary-appui/modules/diary/services"

	"github.com/gin-gonic/gin"
)

func World(ctx *gin.Context) {
	hasError := world(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
	}
	diary, err := diaryService.World(ctx)
	if err != nil {
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx, diary)
}

func world(ctx *gin.Context) bool {
	var inputs diaryModel.DiaryShowPars
	err := ctx.ShouldBind(inputs)
	if err != nil {
		return true
	}
	return false
}
