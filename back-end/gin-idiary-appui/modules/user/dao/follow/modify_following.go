/*
 * @Author: liziwei01
 * @Date: 2022-04-16 17:54:18
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 15:43:21
 * @Description: file content
 */
package follow

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
)

func ModifyFollowing(ctx context.Context, userID int64, followingList string) error {
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return err
	}

	tableName := USER_FOLLOWING_TABLE

	data := make([]map[string]interface{}, 0)
	update := map[string]interface{}{
		"user_id":        userID,
		"following_list": followingList,
	}
	data = append(data, update)

	_, err = client.InsertOnDuplicate(ctx, tableName, data, update)
	if err != nil {
		return err
	}

	return nil
}
