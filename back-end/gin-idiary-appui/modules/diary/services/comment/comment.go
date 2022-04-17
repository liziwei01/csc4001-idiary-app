/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:56:00
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 15:01:43
 * @Description: file content
 */
package comment

import (
	"context"
	commentData "gin-idiary-appui/modules/diary/data/comment"
	commentModel "gin-idiary-appui/modules/diary/model/comment"
)

func Comment(ctx context.Context, pars commentModel.CommentInfo) error {
	return commentData.Comment(ctx, pars.UserID, pars.DiaryID, pars.Content)
}
