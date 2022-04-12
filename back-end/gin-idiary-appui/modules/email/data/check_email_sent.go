/*
 * @Author: liziwei01
 * @Date: 2022-04-12 13:27:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 13:43:47
 * @Description: file content
 */
package data

import (
	"context"
	"fmt"
	"gin-idiary-appui/library/redis"
	"gin-idiary-appui/modules/email/constant"
)

func EmailSent(ctx context.Context, email string) (bool, error) {
	client, err := redis.GetClient(ctx, constant.SERVICE_CONF_REDIS_IDIARY)
	if err != nil {
		return false, err
	}

	key := getEmailKey(email)

	return client.Exists(ctx, key)
}

func getEmailKey(email string) string {
	return fmt.Sprintf(constant.CACHED_USER_EMAIL, email)
}
