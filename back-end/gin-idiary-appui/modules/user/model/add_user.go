/*
 * @Author: liziwei01
 * @Date: 2022-03-03 19:45:35
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-04 15:59:22
 * @Description: model
 */
package model

type UserRegisterPars struct {
	UserID string `form:"user_id" json:"user_id" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}
