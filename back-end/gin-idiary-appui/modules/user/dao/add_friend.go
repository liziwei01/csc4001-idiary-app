/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 10:54:27
 * @Description: file content
 */
package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
	userModel "gin-idiary-appui/modules/user/model"
)

func AddFriend(ctx context.Context, pars userModel.FriendsPars) error {
	client, err := mysql.GetClient(ctx, constant.MYSQL_DB_IDIARY)
	if err != nil {
		return err
	}

	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"user_id":   pars.UserID,
		"friend_id": pars.FriendID,
	})

	_, err = client.Insert(ctx, "idiary_friends", mapSliceInsertData)
	if err != nil {
		return err
	}

	return nil
}
