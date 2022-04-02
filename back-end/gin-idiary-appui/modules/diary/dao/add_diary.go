package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AddDiary(ctx context.Context, pars diaryModel.DiaryRegisterPars) error {
	client, err := mysql.GetClient(ctx, constant.MYSQL_DB_IDIARY_DIARY)
	if err != nil {
		return err
	}
	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"user_id":   pars.UserID,
		"title":     pars.Title,
		"content":   pars.Content,
		"timestamp": pars.Timestamp,
		"authority": pars.Authority,
		"address":   pars.Address,

		// "nickname": pars.Nickname,
		// "email":    pars.Email,
	})
	_, err = client.Insert(ctx, constant.MYSQL_USER_PRIVATE_TABLE_NAME, mapSliceInsertData)
	if err != nil {
		return err
	}
	return nil
}
