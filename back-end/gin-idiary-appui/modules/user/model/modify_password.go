/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 21:23:02
 * @Description: file content
 */
package model

type UserPars struct {
	UserID      string `form:"user_id" json:"user_id" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required"`
}
