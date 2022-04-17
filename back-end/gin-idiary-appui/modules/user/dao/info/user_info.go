/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:20:00
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:21:31
 * @Description: file content
 */
package info

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
	infoModel "gin-idiary-appui/modules/user/model/info"
)

const (
	//用户隐私数据表
	USER_PRIVATE_INFO_TABLE = "tb_user_private_info"
)

func GetUserInfoByEmail(ctx context.Context, email string) ([]infoModel.UserInfo, error) {
	var userInfo []infoModel.UserInfo

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return make([]infoModel.UserInfo, 0), err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	where := map[string]interface{}{
		"email": email,
	}

	columns := []string{"*"}

	err = client.Query(ctx, tableName, where, columns, &userInfo)
	if err != nil {
		return make([]infoModel.UserInfo, 0), err
	}

	return userInfo, nil
}

func GetUserInfoByUserID(ctx context.Context, userID int64) (infoModel.UserInfo, error) {
	var userInfo infoModel.UserInfo

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	where := map[string]interface{}{
		"user_id": userID,
	}

	columns := []string{"*"}

	err = client.Query(ctx, tableName, where, columns, &userInfo)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	return userInfo, nil
}
