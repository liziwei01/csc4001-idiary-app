/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:13:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-18 20:19:37
 * @Description: file content
 */
package info

type UserInfoRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type UserInfo struct {
	UserID   string `json:"user_id" ddb:"user_id"`
	Nickname string `json:"nickname" ddb:"nickname"`
	Profile  string `json:"profile" ddb:"profile"`
	Password string `json:"password" ddb:"password"`
}
