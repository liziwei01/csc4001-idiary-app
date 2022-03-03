/*
 * @Author: liziwei01
 * @Date: 2022-03-03 19:45:35
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-03 20:10:24
 * @Description: file content
 */
package model

type UserRegisterPars struct {
	UserID   string `form:"user_id" json:"user_id" binding:"required"`
	// Nickname string 
	// Email    string
}
