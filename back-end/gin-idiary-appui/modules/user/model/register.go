/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 20:28:21
 * @Description: file content
 */
package model

type RegisterPars struct {
	Email            string `form:"email" json:"email" binding:"required"`
	VerificationCode string `form:"verification_code" json:"verification_code" binding:"required"`
	Username         string `form:"username" json:"username" binding:"required"`
	Password         string `form:"password" json:"password" binding:"required"`
}
