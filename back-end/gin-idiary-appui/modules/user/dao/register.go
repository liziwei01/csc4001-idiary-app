package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
	userModel "gin-idiary-appui/modules/user/model"
)

func Register(ctx context.Context, pars userModel.RegisterPars) error {
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"password": pars.Password,
		"nickname": pars.Username,
	})

	_, err = client.Insert(ctx, tableName, mapSliceInsertData)
	if err != nil {
		return err
	}

	return nil
}
