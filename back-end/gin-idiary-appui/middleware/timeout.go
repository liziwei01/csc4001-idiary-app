/*
 * @Author: liziwei01
 * @Date: 2022-03-10 16:16:18
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-10 16:36:27
 * @Description: file content
 */
package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// timeout middleware wraps the request context with a timeout
func ReadTimeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {

			// wrap the request context with a timeout
			ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

			defer func() {
				// check if context timeout was reached
				if ctx.Err() == context.DeadlineExceeded {

					// write response and abort the request
					c.Writer.WriteHeader(http.StatusGatewayTimeout)
					c.Abort()
				}

				//cancel to clear resources after finished
				cancel()
			}()

			// replace request with context wrapped request
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			c.Next()
		}
	}
}

func WriteTimeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {

			// wrap the request context with a timeout
			ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

			defer func() {
				// check if context timeout was reached
				if ctx.Err() == context.DeadlineExceeded {

					// write response and abort the request
					c.Writer.WriteHeader(http.StatusGatewayTimeout)
					c.Abort()
				}

				//cancel to clear resources after finished
				cancel()
			}()

			// replace request with context wrapped request
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			c.Next()
		}
	}
}
