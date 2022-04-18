/*
 * @Author: liziwei01
 * @Date: 2022-04-18 17:27:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-18 21:51:03
 * @Description: file content
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
	router.POST("/api/user/modifyProfile", userControllers.ModifyProfile)
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", userControllers.Register)
		userGroup.POST("/login", userControllers.Login)
		userGroup.POST("/modifyPassword", userControllers.ModifyPassword)
		userGroup.POST("/getUserInfo", userControllers.Info)
		userGroup.POST("/getAllUserInfo", userControllers.AllInfo)

		followGroup := userGroup.Group("/follow")
		{
			followGroup.POST("/follow", followControllers.Follow)
			followGroup.POST("/followings", followControllers.Followings)
			followGroup.POST("/followers", followControllers.Followers)
		}
	}
}
