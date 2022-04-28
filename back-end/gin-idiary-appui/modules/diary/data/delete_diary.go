/*
 * @Author: liziwei01
 * @Date: 2022-04-09 23:52:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:53:06
 * @Description: file content
 */
package data

import (
	"context"
	diaryDao "gin-idiary-appui/modules/diary/dao"
)

func DeleteDiary(ctx context.Context, diaryID int64) error {
	return diaryDao.DeleteDiary(ctx, diaryID)
}
