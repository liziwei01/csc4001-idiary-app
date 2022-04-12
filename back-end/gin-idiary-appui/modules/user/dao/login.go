/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 10:51:23
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
	userModel "gin-idiary-appui/modules/user/model"
)

func Login(ctx context.Context, pars userModel.LoginPars) (string, error) {
	var password string

	// 数据库名字，之后替换
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY)
	if err != nil {
		return "", err
	}

	where := map[string]interface{}{
		"user_id": pars.UserID,
	}

	columns := []string{"password"}

	// "idiary_diary"是表名，之后替换
	// diary是一个list
	err = client.Query(ctx, "idiary_user", where, columns, &password)
	if err != nil {
		return "", err
	}

	return password, nil // 密码验证
}
