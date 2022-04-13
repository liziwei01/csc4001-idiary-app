/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:14:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-14 00:09:45
 * @Description: file content
 */
package controllers

import (
	"fmt"
	"gin-idiary-appui/library/logit"
	"gin-idiary-appui/library/response"
	uploadModel "gin-idiary-appui/modules/upload/model"
	uploadService "gin-idiary-appui/modules/upload/services"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(ctx *gin.Context) {
	inputs, hasError := getUploadImagePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	src, err := uploadService.UploadImage(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, src)
}

func getUploadImagePars(ctx *gin.Context) (uploadModel.UploadPars, bool) {
	var inputs uploadModel.UploadPars

	startFormat := time.Now()
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	costFormat := time.Since(startFormat)

	logit.Logger.Info(fmt.Sprintf("get upload par from request cost=[%s], file bytes is %d", costFormat, inputs.FileHeader.Size))

	return inputs, false
}
