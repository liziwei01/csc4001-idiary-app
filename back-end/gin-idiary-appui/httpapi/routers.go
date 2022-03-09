/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:46
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-09 20:27:25
 * @Description: 路由分发
 */

package httpapi

import (
	"net/http"

	userRouters "gin-idiary-appui/modules/user/routers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: start http server and start listening
 * @param {*}
 * @return {*}
 */
func InitRouters(router *gin.Engine) {
	// router.Use(middleware.CheckTokenMiddleware(), middleware.GetFrequencyControlMiddleware(), middleware.PostFrequencyControlMiddleware(), middleware.MailFrequencyControlMiddleware())
	// init routers
	userRouters.Init(router)

	// safe router
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello! THis is iDiary. Welcome to our offical website!")
	})
}
