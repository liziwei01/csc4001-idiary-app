/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:16:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 19:51:00
 * @Description: 路由
 */

package routers

import (
	userControllers "gin-idiary-appui/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", userControllers.Login)
		userGroup.POST("/register", userControllers.Register)
		userGroup.POST("/addFriend", userControllers.AddFriend)
		userGroup.POST("/modifyPassword", userControllers.ModifyPassword)
	}
}
