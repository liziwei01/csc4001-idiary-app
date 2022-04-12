/*
 * @Author: liziwei01
 * @Date: 2022-04-12 21:54:39
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 21:55:41
 * @Description: file content
 */
package middleware

import "github.com/gin-gonic/gin"

func CrossRegionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
