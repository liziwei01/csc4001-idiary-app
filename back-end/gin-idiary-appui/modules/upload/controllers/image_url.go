/*
 * @Author: liziwei01
 * @Date: 2022-04-12 15:31:28
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 15:35:08
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	uploadModel "gin-idiary-appui/modules/upload/model"
	uploadService "gin-idiary-appui/modules/upload/services"

	"github.com/gin-gonic/gin"
)

func GetImageURL(ctx *gin.Context) {
	inputs, hasError := getImageURLPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
	}
	src, err := uploadService.GetImageURL(ctx, inputs.FileName)
	if err != nil {
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx, src)
}

func getImageURLPars(ctx *gin.Context) (uploadModel.ImageURL, bool) {
	var inputs uploadModel.ImageURL

	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
