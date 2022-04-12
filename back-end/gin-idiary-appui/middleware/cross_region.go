/*
 * @Author: liziwei01
 * @Date: 2022-04-12 21:54:39
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 22:30:14
 * @Description: file content
 */
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrossRegionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Access-Control-Max-Age", "1728000")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok")
		}
		c.Next()
	}
}
