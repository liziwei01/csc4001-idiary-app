/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 11:02:31
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
	userModel "gin-idiary-appui/modules/user/model"
)

func ModifyPassword(ctx context.Context, pars userModel.UserPars) error {
	// 数据库名字，之后替换
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY)

	where := map[string]interface{}{
		"user_id": pars.UserID,
	}

	update_password := map[string]interface{}{
		"password": pars.NewPassword,
	}

	// "idiary_diary"是表名，之后替换
	// diary是一个list
	_, err = client.Update(ctx, "idiary_user", where, update_password)
	if err != nil {
		return err
	}

	return nil
}
