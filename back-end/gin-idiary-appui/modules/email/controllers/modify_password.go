package controllers

import (
	"gin-idiary-appui/library/response"
	emailModel "gin-idiary-appui/modules/email/model"
	emailService "gin-idiary-appui/modules/email/services"

	"github.com/gin-gonic/gin"
)

func ModifyPassword(ctx *gin.Context, password string) {
	inputs, hasError := getUserPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
	}
	err := emailService.ModifyPassword(ctx, inputs, password)
	if err != nil {
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx)
}

func getUserPars(ctx *gin.Context) (emailModel.UserPars, bool) {
	var inputs emailModel.UserPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
