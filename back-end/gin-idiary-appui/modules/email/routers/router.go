/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:16:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-10 17:22:51
 * @Description: 路由
 */

package routers

import (
	userControllers "gin-idiary-appui/modules/email/controllers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.Engine) {
	userGroup := router.Group("/email")
	{
		userGroup.GET("/addUser", userControllers.AddUser)
		userGroup.GET("/modifyPassword", userControllers.ModifyPassword)
	}
}
