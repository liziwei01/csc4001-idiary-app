/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:14:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-14 00:09:36
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	emailModel "gin-idiary-appui/modules/email/model"
	emailService "gin-idiary-appui/modules/email/services"

	"github.com/gin-gonic/gin"
)

func VerificationCode(ctx *gin.Context) {
	inputs, hasError := getVerificationCodePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := emailService.VerificationCode(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getVerificationCodePars(ctx *gin.Context) (emailModel.EmailPars, bool) {
	var inputs emailModel.EmailPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
