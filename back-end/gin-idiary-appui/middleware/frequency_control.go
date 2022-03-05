/*
 * @Author: liziwei01
 * @Date: 2022-03-04 21:44:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-05 15:12:56
 * @Description: 频控中间件
 */
package middleware

import (
	"github.com/gin-gonic/gin"
)

func GetFrequencyControlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.FreqControl.Enable {
			// 不限制.
			c.Next()
		}
	}
}

func PostFrequencyControlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.FreqControl.Enable {
			// 不限制.
			c.Next()
		}
	}
}

func MailFrequencyControlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.FreqControl.Enable {
			// 不限制.
			c.Next()
		}
	}
}
