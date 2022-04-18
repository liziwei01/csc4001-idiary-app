package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
	userModel "gin-idiary-appui/modules/user/model"
)

func AllUsers(ctx context.Context, pars userModel.UserListRequestPars) ([]userModel.UserInfo, error) {
	var user []userModel.UserInfo

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return nil, err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	where := map[string]interface{}{
		"_orderby": "db_time desc",
		"_limit":   []uint{pars.PageIndex, pars.PageLength},
	}

	columns := []string{"*"}

	err = client.Query(ctx, tableName, where, columns, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func AllUserCount(ctx context.Context, pars userModel.UserListRequestPars) (int64, error) {
	var count = make([]struct {
		Count int64 `ddb:"count"`
	}, 1)

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return 0, err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	where := map[string]interface{}{
		"user_id !=": 0,
	}

	columns := []string{"count(1) as count"}

	err = client.Query(ctx, tableName, where, columns, &count)
	if err != nil {
		return 0, err
	}

	return count[0].Count, nil
}
