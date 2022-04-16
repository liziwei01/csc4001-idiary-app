/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 21:02:24
 * @Description: file content
 */
package routers

import (
	diaryControllers "gin-idiary-appui/modules/diary/controllers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.Engine) {
	userGroup := router.Group("/api/diary")
	{
		userGroup.GET("/world", diaryControllers.AllDiary)
		userGroup.GET("/friend", diaryControllers.FriendDiary)
		userGroup.POST("/addDiary", diaryControllers.AddDiary)
	}
}
