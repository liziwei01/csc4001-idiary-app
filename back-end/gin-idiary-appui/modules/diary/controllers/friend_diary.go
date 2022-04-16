/*
 * @Author: liziwei01
 * @Date: 2022-04-16 20:00:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 20:00:14
 * @Description: file content
 */
package controllers

import (
	"gin-idiary-appui/library/response"
	diaryService "gin-idiary-appui/modules/diary/services"

	"github.com/gin-gonic/gin"
)

func FriendDiary(ctx *gin.Context) {
	inputs, hasError := getAllDiaryPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	diaries, count, err := diaryService.AllDiary(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, gin.H{
		"diaries": diaries,
		"count":   count,
	})
}
