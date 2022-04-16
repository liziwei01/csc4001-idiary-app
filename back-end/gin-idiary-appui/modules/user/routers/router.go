/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:16:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 18:32:23
 * @Description: 路由
 */

package routers

import (
	userControllers "gin-idiary-appui/modules/user/controllers"
	followControllers "gin-idiary-appui/modules/user/controllers/follow"

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

		followGroup := userGroup.Group("/follow")
		{
			followGroup.POST("/follow", followControllers.Follow)
			followGroup.POST("/followings", followControllers.Followings)
			followGroup.POST("/followers", followControllers.Followers)
		}
		
		userGroup.POST("/modifyPassword", userControllers.ModifyPassword)
	}
}
