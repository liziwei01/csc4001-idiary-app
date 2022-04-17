/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 16:20:49
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func GetDiaryByUserID(ctx context.Context, pars diaryModel.DiaryListRequestPars) ([]diaryModel.DiaryInfo, error) {
	var diary []diaryModel.DiaryInfo

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_FEED)
	if err != nil {
		return nil, err
	}

	tableName := DIARY_FEED_TABLE

	where := map[string]interface{}{
		"_orderby": "db_time desc",
		"_limit":   []uint{pars.PageIndex, pars.PageLength},
		"user_id":  pars.UserID,
	}

	columns := []string{"*"}

	err = client.Query(ctx, tableName, where, columns, &diary)
	if err != nil {
		return nil, err
	}

	return diary, nil
}

func GetDiaryByUserIDCount(ctx context.Context, pars diaryModel.DiaryListRequestPars) (int64, error) {
	var count = make([]struct {
		Count int64 `ddb:"count"`
	}, 1)

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_FEED)
	if err != nil {
		return 0, err
	}

	tableName := DIARY_FEED_TABLE

	where := map[string]interface{}{
		"user_id": pars.UserID,
	}

	columns := []string{"count(1) as count"}

	err = client.Query(ctx, tableName, where, columns, &count)
	if err != nil {
		return 0, err
	}

	return count[0].Count, nil
}
