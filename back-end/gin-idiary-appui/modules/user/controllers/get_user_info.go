/*
 * @Author: liziwei01
 * @Date: 2022-04-18 20:12:54
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-18 21:50:01
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	infoModel "gin-idiary-appui/modules/user/model/info"

	infoService "gin-idiary-appui/modules/user/services/info"

	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	inputs, hasError := getInfoPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	info, err := infoService.Info(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}

	response.StdSuccess(ctx, gin.H{
		"user_id":  info.UserID,
		"nickname": info.Nickname,
		"profile":  info.Profile,
	})
}

func AllInfo(ctx *gin.Context) {
	inputs, hasError := getInfoPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	infos, err := infoService.AllInfo(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}

	response.StdSuccess(ctx, infos)
}

func getInfoPars(ctx *gin.Context) (infoModel.UserInfoRequest, bool) {
	var inputs infoModel.UserInfoRequest
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
