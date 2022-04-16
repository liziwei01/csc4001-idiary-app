/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:35:42
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
	diaryModel "gin-idiary-appui/modules/diary/model"
	"time"
)

const (
	// 日记投稿总表
	DIARY_FEED_TABLE = "tb_user_diary_feed"
)

func AddDiary(ctx context.Context, pars diaryModel.DiaryInfo) error {
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_FEED)
	if err != nil {
		return err
	}

	tableName := DIARY_FEED_TABLE

	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"user_id":    pars.UserID,
		"content":    pars.Content,
		"image_list": pars.ImageList,
		"device":     pars.Device,
		"db_time":    time.Now().Unix(),
		"authority":  pars.Authority,
		"address":    pars.Address,
	})

	_, err = client.Insert(ctx, tableName, mapSliceInsertData)
	if err != nil {
		return err
	}

	return nil
}
