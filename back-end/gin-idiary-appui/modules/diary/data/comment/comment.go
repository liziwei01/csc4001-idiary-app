/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:57:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 15:01:00
 * @Description: file content
 */
package comment

import (
	"context"
	commentDao "gin-idiary-appui/modules/diary/dao/comment"
)

func Comment(ctx context.Context, userID, diaryID int64, content string) error {
	return commentDao.Comment(ctx, userID, diaryID, content)
}
