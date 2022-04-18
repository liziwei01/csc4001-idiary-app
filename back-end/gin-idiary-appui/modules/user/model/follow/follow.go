/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-18 20:07:12
 * @Description: file content
 */
package model

type FollowPars struct {
	UserID      int64 `form:"user_id" json:"user_id" binding:"required"`
	FollowingID int64 `form:"following_id" json:"following_id" binding:"required"`
}

type FollowingPars struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type FollowerPars struct {
	Email string `form:"email" json:"email" binding:"required"`
}
