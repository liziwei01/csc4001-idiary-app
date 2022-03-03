/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:46
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-03 19:23:40
 * @Description: file content
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
func InitRouters() *gin.Engine {
	// init routers
	router := gin.Default()
	userRouters.Init(router)

	// safe router
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello! THis is iDiary. Welcome to our website!")
	})
	return router
}
