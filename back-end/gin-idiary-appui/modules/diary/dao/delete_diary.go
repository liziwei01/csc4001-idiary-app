/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-25 21:27:57
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
)

func DeleteDiary(ctx context.Context, diaryID int64) error {
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_FEED)
	if err != nil {
		return err
	}

	tableName := DIARY_FEED_TABLE

	where := map[string]interface{}{
		"diary_id": diaryID,
	}

	update := map[string]interface{}{
		"delete_status": 1,
	}

	_, err = client.Update(ctx, tableName, where, update)
	if err != nil {
		return err
	}

	return nil
}
