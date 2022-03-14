/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:16:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-10 17:23:01
 * @Description: 路由
 */

package routers

import (
	userControllers "gin-idiary-appui/modules/diary/controllers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.Engine) {
	userGroup := router.Group("/diary")
	{
		userGroup.GET("/addUser", userControllers.AddUser)
		userGroup.GET("/addDiary", userControllers.AddDiary)
		userGroup.GET("/allDiary", userControllers.AllDiary)
		userGroup.GET("/modifyDiary", userControllers.ModifyDiary)
		userGroup.GET("/deleteDiary", userControllers.DeleteDiary)
		userGroup.GET("/searchDiary", userControllers.SearchDiary)
		userGroup.GET("setPrivacy", userControllers.SetPrivacy)
	}
}
