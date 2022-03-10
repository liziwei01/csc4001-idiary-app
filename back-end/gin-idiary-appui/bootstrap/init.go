/*
 * @Author: liziwei01
 * @Date: 2022-03-04 22:06:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-10 16:39:05
 * @Description: file content
 */
package bootstrap

import (
	"context"
	"gin-idiary-appui/library/logit"
	"gin-idiary-appui/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMust(ctx context.Context) {
	InitLog(ctx)
	InitMiddleware(ctx)
}

func InitLog(ctx context.Context) {
	logit.Init(ctx, "gin-idiary-appui")
}

func InitRedis() {

}

func InitMiddleware(ctx context.Context) {
	middleware.Init(ctx)
}

// InitHandler 获取http handler
func InitHandler(app *AppServer) *gin.Engine {
	gin.SetMode(app.Config.RunMode)
	handler := gin.New()
	// 注册log recover中间件
	handler.Use(gin.Logger(), gin.Recovery())
	// 注册超时中间件
	handler.Use(middleware.ReadTimeoutMiddleware(time.Duration(app.Config.HTTPServer.ReadTimeout)), middleware.WriteTimeoutMiddleware(time.Duration(app.Config.HTTPServer.WriteTimeout)))
	return handler
}
