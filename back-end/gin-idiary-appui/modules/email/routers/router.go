package routers

import (
	emailControllers "gin-idiary-appui/modules/email/controllers"

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
		userGroup.GET("/modifyPassword", emailControllers.ModifyPassword)
	}
}
