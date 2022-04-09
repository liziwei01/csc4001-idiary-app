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
	userGroup := router.Group("/diary")
	{
		userGroup.GET("/allDiary", diaryControllers.AllDiary)
		userGroup.GET("/addDiary", diaryControllers.AddDiary)
	}
}
