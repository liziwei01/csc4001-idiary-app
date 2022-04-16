/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 18:45:35
 * @Description: file content
 */
package model

type DiaryPars struct {
	UserID    string `form:"user_id" json:"user_id" binding:"required"`
	DiaryID   string `form:"diary_id" json:"diary_id"`
	Title     string `form:"title" json:"title"`
	Content   string `form:"content" json:"content" binding:"required"`
	ImageList string `form:"image_list" json:"image_list"`
	DBTime    int64  `form:"timestamp" json:"timestamp"`
	Authority string `form:"authority" json:"authority"`
	Device    string `form:"device" json:"device"`
	Address   string `form:"address" json:"address"`
}
