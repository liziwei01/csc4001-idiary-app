/*
 * @Author: liziwei01
 * @Date: 2022-04-24 20:41:04
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-25 09:52:21
 * @Description: file content
 */
package comment

import (
	"context"
	"fmt"
	"gin-idiary-appui/library/redis"
	"gin-idiary-appui/modules/diary/constant"
	"time"

	"github.com/gogf/gf/util/gconv"
)

const (
	SECONDS_PER_YEAR                   = 60 * 60 * 24 * 365
	USER_LIKE_DIARY_EXPIRE_TIME        = SECONDS_PER_YEAR
	USER_DIARY_LIKED_COUNT_EXPIRE_TIME = SECONDS_PER_YEAR * 100
)

func Like(ctx context.Context, userID, diaryID int64) error {
	client, err := redis.GetClient(ctx, constant.SERVICE_CONF_REDIS_IDIARY)
	if err != nil {
		return err
	}

	key := getUserLikeDiaryKey(userID, diaryID)

	err = client.Set(ctx, key, "1", getUserLikeDiaryExpireNanoSecond())
	if err != nil {
		return err
	}

	key = getDiaryLikeCountKey(diaryID)

	likeCount, err := DiaryLikeCount(ctx, diaryID)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, gconv.String(likeCount+1), getDiaryLikeCountExpireNanoSecond())
}

func UnLike(ctx context.Context, userID, diaryID int64) error {
	client, err := redis.GetClient(ctx, constant.SERVICE_CONF_REDIS_IDIARY)
	if err != nil {
		return err
	}

	key := getUserLikeDiaryKey(userID, diaryID)

	err = client.Del(ctx, key)
	if err != nil {
		return err
	}

	key = getDiaryLikeCountKey(diaryID)

	likeCount, err := DiaryLikeCount(ctx, diaryID)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, gconv.String(likeCount-1), getDiaryLikeCountExpireNanoSecond())
}

func getUserLikeDiaryKey(userID int64, diaryID int64) string {
	return fmt.Sprintf(constant.CACHED_USER_LIKE_DIARY, userID, diaryID)
}

func getDiaryLikeCountKey(diaryID int64) string {
	return fmt.Sprintf(constant.CACHED_DIARY_LIKED_COUNT, diaryID)
}

// 用户点赞过期时间: 一年
func getUserLikeDiaryExpireNanoSecond() time.Duration {
	return time.Second * time.Duration(USER_LIKE_DIARY_EXPIRE_TIME)
}

// 日记点赞数过期时间: 一百年
func getDiaryLikeCountExpireNanoSecond() time.Duration {
	return time.Second * time.Duration(USER_DIARY_LIKED_COUNT_EXPIRE_TIME)
}
