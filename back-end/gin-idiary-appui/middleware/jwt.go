/*
 * @Author: liziwei01
 * @Date: 2022-04-13 23:08:22
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-13 23:45:34
 * @Description: file content
 */
package middleware

import (
	"gin-idiary-appui/library/jwtoken"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT gin middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// skip /user/login & /user/register
		if c.Request.URL.Path == "/user/login" || c.Request.URL.Path == "/user/register" {
			c.Next()
			return
		}

		// Get the JWT token from the header
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"code":    1,
				"message": "请求头中未携带token",
			})
			c.Abort()
			return
		}
		// Parse the token
		jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtoken.JWTSecret, nil
		})
		if err != nil {
			c.JSON(401, gin.H{
				"code":    1,
				"message": "token错误",
			})
			c.Abort()
			return
		}
		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
			c.Set("claims", claims)
		} else {
			c.JSON(401, gin.H{
				"code":    1,
				"message": "无效的token",
			})
			c.Abort()
			return
		}
	}
}
