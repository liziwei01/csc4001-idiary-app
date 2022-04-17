/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:03:53
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:54:00
 * @Description: file content
 */
package comment

type CommentInfo struct {
	CommentID    int64  `form:"comment_id" json:"comment_id" ddb:"comment_id"`
	DiaryID      int64  `form:"diary_id" json:"diary_id" ddb:"diary_id" binding:"required"`
	UserID       int64  `form:"user_id" json:"user_id" ddb:"user_id" binding:"required"`
	Nickname     string `form:"nickname" json:"nickname" ddb:"nickname"`
	Profile      string `form:"profile" json:"profile" ddb:"profile"`
	Content      string `form:"content" json:"content" ddb:"content" binding:"required"`
	Device       string `form:"device" json:"device" ddb:"device"`
	DBTime       int64  `form:"db_time" json:"db_time" ddb:"db_time"`
	Address      string `form:"address" json:"address" ddb:"address"`
	VoteCount    int64  `form:"vote_count" json:"vote_count" ddb:"vote_count"`
	DislikeCount int64  `form:"dislike_count" json:"dislike_count" ddb:"dislike_count"`
	ReportCount  int64  `form:"report_count" json:"report_count" ddb:"report_count"`
	DeleteStatus uint8  `form:"delete_status" json:"delete_status" ddb:"delete_status"`
}
