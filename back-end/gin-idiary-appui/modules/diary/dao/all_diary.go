/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 20:15:17
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryShowPars) ([]diaryModel.DiaryInfo, error) {

	var diary []diaryModel.DiaryInfo

	// 数据库名字，之后替换
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY)
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
