/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:15:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 13:23:34
 * @Description: file content
 */
package model

type EmailPars struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}
