package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AddDiary(ctx context.Context, pars diaryModel.DiaryRegisterPars) error {
	client, err := mysql.GetClient(ctx, "idiary")
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

		// "nickname": pars.Nickname,
		// "email":    pars.Email,
	})
	_, err = client.Insert(ctx, "idiary_diary", mapSliceInsertData)
	if err != nil {
		return err
	}
	return nil
}
