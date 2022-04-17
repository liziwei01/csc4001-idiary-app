/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 15:24:21
 * @Description: file content
 */
package follow

import (
	"context"
	"encoding/json"
	"gin-idiary-appui/library/logit"
	followDao "gin-idiary-appui/modules/user/dao/follow"

	"github.com/gogf/gf/util/gconv"
)

func Follow(ctx context.Context, userID, followingID int64) error {
	var followingList = make([]int64, 0)
	var followerList = make([]int64, 0)

	// 1 拿到关注列表
	followings, err := followDao.GetFollowing(ctx, userID)
	if err != nil {
		return err
	}

	// 2 添加关注
	if followings != "" {
		err = json.Unmarshal([]byte(followings), &followingList)
		if err != nil {
			logit.Logger.Error("json.Unmarshal error: #+v", err)
			return err
		}
	}

	followingList = append(followingList, followingID)
	followingsBytes, err := json.Marshal(followingList)
	if err != nil {
		logit.Logger.Error("json.Marshal error: #+v", err)
		return err
	}

	// 3 更新关注列表
	err = followDao.ModifyFollowing(ctx, userID, gconv.String(followingsBytes))
	if err != nil {
		return err
	}

	// 4 拿到被关注者的粉丝列表
	followers, err := followDao.GetFollower(ctx, followingID)
	if err != nil {
		return err
	}

	// 5 添加粉丝
	if followers != "" {
		err = json.Unmarshal([]byte(followers), &followerList)
		if err != nil {
			logit.Logger.Error("json.Unmarshal error: #+v", err)
			return err
		}
	}

	followerList = append(followerList, userID)
	followersBytes, err := json.Marshal(followerList)
	if err != nil {
		logit.Logger.Error("json.Marshal error: #+v", err)
		return err
	}

	// 6 更新粉丝列表
	err = followDao.ModifyFollower(ctx, followingID, gconv.String(followersBytes))
	if err != nil {
		return err
	}

	return nil
}
