/*
 * @Author: liziwei01
 * @Date: 2022-04-24 20:36:52
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-25 09:49:19
 * @Description: file content
 */
package comment

type DiaryLikeRequestPars struct {
	DiaryID      int64 `form:"diary_id" json:"diary_id" ddb:"diary_id" binding:"required"`
	UserID       int64 `form:"user_id" json:"user_id" ddb:"user_id" binding:"required"`
	ShouldUnlike bool  `form:"should_unlike" json:"should_unlike"`
}
