/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 11:01:21
 * @Description: file content
 */
package model

type UserPars struct {
	UserID      string `form:"user_id" json:"user_id" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	NewPassword string `form:"newPassword" json:"newPassword" binding:"required"`
	Nickname    string `form:"nickname" json:"nickname" binding:"required"`
	City        string `form:"city" json:"city" binding:"required"`
	Picture     string `form:"picture" json:"picture" binding:"required"`
}
