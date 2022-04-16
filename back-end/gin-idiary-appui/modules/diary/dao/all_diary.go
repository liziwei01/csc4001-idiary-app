/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 18:46:06
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryListRequestPars) ([]diaryModel.DiaryInfo, error) {
	var diary []diaryModel.DiaryInfo

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_FEED)
	if err != nil {
		return nil, err
	}

	where := map[string]interface{}{
		"user_id": pars.UserID,
	}

	columns := []string{"*"}

	// "idiary_diary"是表名，之后替换
	// diary是一个list
	_ = client.Query(ctx, "idiary_diary", where, columns, &diary)
	if err != nil {
		return nil, err
	}

	return diary, nil
}
