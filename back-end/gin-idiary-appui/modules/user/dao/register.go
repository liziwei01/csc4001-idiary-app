/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 10:50:51
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
	userModel "gin-idiary-appui/modules/user/model"
)

func Register(ctx context.Context, pars userModel.RegisterPars) error {
	client, err := mysql.GetClient(ctx, constant.MYSQL_DB_IDIARY)
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
	})

	_, err = client.Insert(ctx, "idiary_user", mapSliceInsertData)
	if err != nil {
		return err
	}

	return nil
}
