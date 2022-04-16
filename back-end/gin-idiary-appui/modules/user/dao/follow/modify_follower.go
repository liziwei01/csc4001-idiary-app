/*
 * @Author: liziwei01
 * @Date: 2022-04-16 17:54:18
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 18:28:21
 * @Description: file content
 */
package follow

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
)

func ModifyFollower(ctx context.Context, userID int64, followerList string) error {
	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return err
	}

	tableName := USER_FOLLOWER_TABLE

	where := map[string]interface{}{
		"user_id": userID,
	}

	update := map[string]interface{}{
		"follower_list": followerList,
	}

	_, err = client.Update(ctx, tableName, where, update)
	if err != nil {
		return err
	}

	return nil
}
