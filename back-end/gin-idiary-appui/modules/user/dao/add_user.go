/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:19:03
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-09 20:34:03
 * @Description: 数据库操作
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
	userModel "gin-idiary-appui/modules/user/model"
)

func AddUser(ctx context.Context, pars userModel.UserRegisterPars) error {
	client, err := mysql.GetMysqlClient(ctx, constant.MYSQL_DB_IDIARY_USER)
	if err != nil {
		return err
	}
	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"user_id": pars.UserID,
		// "nickname": pars.Nickname,
		// "email":    pars.Email,
	})
	_, err = client.Insert(ctx, constant.MYSQL_USER_PRIVATE_TABLE_NAME, mapSliceInsertData)
	if err != nil {
		return err
	}
	return nil
}
