package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	userModel "gin-idiary-appui/modules/user/model"
)

func Register(ctx context.Context, pars userModel.RegisterPars) error {
	client, err := mysql.GetClient(ctx, "idiary")
	if err != nil {
		return err
	}
	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"user_id":  pars.UserID,
		"password": pars.Password,
		"city":     pars.City,
		"nickname": pars.Nickname,
		"picture":  pars.Picture,

		// "nickname": pars.Nickname,
		// "email":    pars.Email,
	})
	_, err = client.Insert(ctx, "idiary_user", mapSliceInsertData)
	if err != nil {
		return err
	}
	return nil
}
