/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 16:09:41
 * @Description: file content
 */
package model

import commentModel "gin-idiary-appui/modules/diary/model/comment"

type DiaryListRequestPars struct {
	UserID     int64 `form:"user_id" json:"user_id" binding:"required"`
	PageIndex  uint  `form:"page_index" json:"page_index"`
	PageLength uint  `form:"page_length" json:"page_length"`
}

type FriendDiaryListRequestPars struct {
	UserID     int64   `form:"user_id" json:"user_id" binding:"required"`
	FriendIDs  []int64 `form:"friend_ids" json:"friend_ids"`
	PageIndex  uint    `form:"page_index" json:"page_index"`
	PageLength uint    `form:"page_length" json:"page_length"`
}

type DiaryInfo struct {
	DiaryID      int64                      `form:"diary_id" json:"diary_id" ddb:"diary_id"`
	UserID       int64                      `form:"user_id" json:"user_id" ddb:"user_id" binding:"required"`
	Nickname     string                     `form:"nickname" json:"nickname"`
	UserProfile  string                     `form:"user_profile" json:"user_profile"`
	Title        string                     `form:"title" json:"title" ddb:"title"`
	Content      string                     `form:"content" json:"content" ddb:"content" binding:"required"`
	ImageList    string                     `form:"image_list" json:"image_list" ddb:"image_list"`
	CommentList  []commentModel.CommentInfo `form:"comment_list" json:"comment_list"`
	CommentCount int64                      `form:"comment_count" json:"comment_count"`
	Device       string                     `form:"device" json:"device" ddb:"device"`
	DBTime       int64                      `form:"db_time" json:"db_time" ddb:"db_time"`
	Authority    uint8                      `form:"authority" json:"authority" ddb:"authority"`
	Address      string                     `form:"address" json:"address" ddb:"address"`
	VoteCount    int64                      `form:"vote_count" json:"vote_count" ddb:"vote_count"`
	HasVoted     bool                       `form:"has_voted" json:"has_voted"`
	DislikeCount int64                      `form:"dislike_count" json:"dislike_count" ddb:"dislike_count"`
	ShareCount   int64                      `form:"share_count" json:"share_count" ddb:"share_count"`
	ReportCount  int64                      `form:"report_count" json:"report_count" ddb:"report_count"`
	DeleteStatus bool                       `form:"delete_status" json:"delete_status" ddb:"delete_status"`
	Tags         string                     `form:"tags" json:"tags" ddb:"tags"`
}
