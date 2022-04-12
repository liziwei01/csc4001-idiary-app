/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 10:55:50
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AddDiary(ctx context.Context, pars diaryModel.DiaryRegisterPars) error {
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY)
	if err != nil {
		return err
	}

	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"user_id":   pars.UserID,
		"diary_id":  pars.DiaryID,
		"title":     pars.Title,
		"content":   pars.Content,
		"timestamp": pars.Timestamp,
		"authority": pars.Authority,
		"address":   pars.Address,
	})

	_, err = client.Insert(ctx, "idiary_diary", mapSliceInsertData)
	if err != nil {
		return err
	}

	return nil
}
