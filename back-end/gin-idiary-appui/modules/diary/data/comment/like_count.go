/*
 * @Author: liziwei01
 * @Date: 2022-04-24 21:39:25
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-24 21:39:26
 * @Description: file content
 */
package comment

import (
	"context"
	"gin-idiary-appui/library/redis"
	"gin-idiary-appui/modules/diary/constant"

	"github.com/gogf/gf/util/gconv"
)

func DiaryLikeCount(ctx context.Context, diaryID int64) (int64, error) {
	var likeCount int64 = 0

	client, err := redis.GetClient(ctx, constant.SERVICE_CONF_REDIS_IDIARY)
	if err != nil {
		return 0, err
	}

	key := getDiaryLikeCountKey(diaryID)

	exists, err := client.Exists(ctx, key)
	if err != nil {
		return 0, err
	}

	if exists {
		likeCountStr, err := client.Get(ctx, key)
		if err != nil {
			return 0, err
		}
		likeCount = gconv.Int64(likeCountStr)
	}

	return likeCount, nil
}
