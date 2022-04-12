/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:46
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 21:51:29
 * @Description: 路由分发
 */

package httpapi

import (
	"net/http"

	"gin-idiary-appui/middleware"
	diaryRouters "gin-idiary-appui/modules/diary/routers"
	emailRouters "gin-idiary-appui/modules/email/routers"
	uploadRouters "gin-idiary-appui/modules/upload/routers"
	userRouters "gin-idiary-appui/modules/user/routers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: start http server and start listening
 * @param {*}
 * @return {*}
 */
func InitRouters(router *gin.Engine) {
	//暂时解决跨域问题
	router.Use(middleware.CrossRegionMiddleware())
	// router.Use(middleware.CheckTokenMiddleware(), middleware.GetFrequencyControlMiddleware(), middleware.PostFrequencyControlMiddleware(), middleware.MailFrequencyControlMiddleware())
	// init routers
	userRouters.Init(router)
	diaryRouters.Init(router)
	emailRouters.Init(router)
	uploadRouters.Init(router)

	// safe router
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello! THis is iDiary. Welcome to our offical website!")
	})
}
