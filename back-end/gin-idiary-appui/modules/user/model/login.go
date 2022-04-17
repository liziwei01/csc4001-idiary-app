/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 23:45:44
 * @Description: file content
 */
package model

type LoginPars struct {
	Email    string `form:"user_id" json:"user_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
