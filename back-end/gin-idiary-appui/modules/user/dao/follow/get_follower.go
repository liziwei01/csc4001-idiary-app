/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 18:28:12
 * @Description: file content
 */
package follow

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
)

const (
	//用户粉丝表
	USER_FOLLOWER_TABLE = "tb_user_followers"
)

func GetFollower(ctx context.Context, userID int64) (string, error) {
	var followers string

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return "", err
	}

	tableName := USER_FOLLOWER_TABLE

	where := map[string]interface{}{
		"user_id": userID,
	}

	columns := []string{"follower_list"}

	err = client.Query(ctx, tableName, where, columns, &followers)
	if err != nil {
		return "", err
	}

	return followers, nil
}
