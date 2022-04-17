/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:58:15
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 15:00:32
 * @Description: file content
 */
package comment

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
	"time"
)

func Comment(ctx context.Context, userID, diaryID int64, content string) error {
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_FEED)
	if err != nil {
		return err
	}

	tableName := DIARY_COMMENT_TABLE

	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"diary_id": diaryID,
		"user_id":  userID,
		"content":  content,
		"db_time":  time.Now().Unix(),
	})

	_, err = client.Insert(ctx, tableName, mapSliceInsertData)
	if err != nil {
		return err
	}

	return nil
}
