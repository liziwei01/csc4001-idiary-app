/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 21:22:14
 * @Description: file content
 */
package dao

import (
	"context"
	"fmt"
	"gin-idiary-appui/library/logit"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
	userModel "gin-idiary-appui/modules/user/model"
)

const (
	//用户隐私数据表
	USER_PRIVATE_INFO_TABLE = "tb_user_private_info"
)

func GetPassword(ctx context.Context, pars userModel.LoginPars) (string, error) {
	var password []string

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return "", err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	where := map[string]interface{}{
		"email": pars.Email,
	}

	columns := []string{"password"}

	err = client.Query(ctx, tableName, where, columns, &password)
	if err != nil {
		return "", err
	}

	if len(password) != 1 {
		err := fmt.Errorf("len(password) = %d", len(password))
		logit.Logger.Warn(err.Error())
		return "", err
	}

	return password[0], nil // 密码验证
}
