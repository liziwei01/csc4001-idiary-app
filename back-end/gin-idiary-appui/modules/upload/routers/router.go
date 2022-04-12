/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 15:30:42
 * @Description: file content
 */
package routers

import (
	"github.com/gin-gonic/gin"

	uploadController "gin-idiary-appui/modules/upload/controllers"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.Engine) {
	uploadGroup := router.Group("/upload")
	{
		uploadGroup.POST("/image", uploadController.UploadImage)
		uploadGroup.GET("/getImageURL", uploadController.GetImageURL)
	}
}
