package dao

import (
	"context"
	"csc4001-idiary-app/back-end/gin-idiary-appui/modules/diary/constant"
	diaryModel "csc4001-idiary-app/back-end/gin-idiary-appui/modules/diary/model"
	"gin-idiary-appui/library/mysql"
	"testing"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryRegisterPars) []diaryModel.DiaryInfo {

	var diary []diaryModel.DiaryInfo

	// 数据库名字，之后替换
	client, err := mysql.GetClient(ctx, "idiary")
	if err != nil {
		return err
	}

	where := map[string]interface{}{
		"user_id": pars.UserID
	}
	columns := []string{"*"}

	// "idiary_diary"是表名，之后替换
	// diary是一个list
	_, err = client.Query(ctx, "idiary_diary", where, columns, &diary)
	if err != nil {
		return err
	}
	return diary
}



