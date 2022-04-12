/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 14:11:39
 * @Description: file content
 */
package routers

import (
	"github.com/gin-gonic/gin"

	emailController "gin-idiary-appui/modules/email/controllers"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.Engine) {
	emailGroup := router.Group("/email")
	{
		emailGroup.GET("/verificationCode", emailController.VerificationCode)
	}
}
