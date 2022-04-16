/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 18:28:30
 * @Description: file content
 */
package follow

import (
	"context"
	"encoding/json"
	"gin-idiary-appui/library/logit"
	followDao "gin-idiary-appui/modules/user/dao/follow"
	userModel "gin-idiary-appui/modules/user/model"

	"github.com/gogf/gf/util/gconv"
)

func Follow(ctx context.Context, pars userModel.FollowPars) error {
	var followingList []int64
	var followerList []int64

	// 1 拿到关注列表
	followings, err := followDao.GetFollowing(ctx, pars.UserID)
	if err != nil {
		return err
	}

	// 2 添加关注
	err = json.Unmarshal([]byte(followings), &followingList)
	if err != nil {
		logit.Logger.Error("json.Unmarshal error: #+v", err)
		return err
	}

	followingList = append(followingList, pars.FollowingID)
	followingsBytes, err := json.Marshal(followingList)
	if err != nil {
		logit.Logger.Error("json.Marshal error: #+v", err)
		return err
	}

	// 3 更新关注列表
	err = followDao.ModifyFollowing(ctx, pars.UserID, gconv.String(followingsBytes))
	if err != nil {
		return err
	}

	// 4 拿到被关注者的粉丝列表
	followers, err := followDao.GetFollower(ctx, pars.FollowingID)
	if err != nil {
		return err
	}

	// 5 添加粉丝
	err = json.Unmarshal([]byte(followers), &followerList)
	if err != nil {
		logit.Logger.Error("json.Unmarshal error: #+v", err)
		return err
	}

	followerList = append(followerList, pars.UserID)
	followersBytes, err := json.Marshal(followerList)
	if err != nil {
		logit.Logger.Error("json.Marshal error: #+v", err)
		return err
	}

	// 6 更新粉丝列表
	err = followDao.ModifyFollower(ctx, pars.FollowingID, gconv.String(followersBytes))
	if err != nil {
		return err
	}

	return nil
}
