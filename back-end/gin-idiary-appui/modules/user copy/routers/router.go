/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:16:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-04 16:00:27
 * @Description: 路由分发
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
		userGroup.GET("/addUser", userControllers.AddUser)
	}
}