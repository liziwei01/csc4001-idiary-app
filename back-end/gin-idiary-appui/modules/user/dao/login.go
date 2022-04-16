/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 17:39:23
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

	// 数据库名字，之后替换
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return "", err
	}

	tableName := USER_PRIVATE_INFO_TABLE

	where := map[string]interface{}{
		"user_id": pars.UserID,
	}

	columns := []string{"password"}

	// "idiary_diary"是表名，之后替换
	// diary是一个list
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
