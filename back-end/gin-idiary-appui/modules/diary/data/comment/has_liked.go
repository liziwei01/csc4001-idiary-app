/*
 * @Author: liziwei01
 * @Date: 2022-04-24 21:39:53
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-24 21:39:54
 * @Description: file content
 */
package comment

import (
	"context"
	"gin-idiary-appui/library/redis"
	"gin-idiary-appui/modules/diary/constant"
)

func DiaryHasLiked(ctx context.Context, userID, diaryID int64) (bool, error) {
	client, err := redis.GetClient(ctx, constant.SERVICE_CONF_REDIS_IDIARY)
	if err != nil {
		return false, err
	}

	key := getUserLikeDiaryKey(userID, diaryID)

	return client.Exists(ctx, key)
}
