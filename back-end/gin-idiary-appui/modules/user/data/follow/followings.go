/*
 * @Author: liziwei01
 * @Date: 2022-04-16 18:24:49
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:16:27
 * @Description: file content
 */
package follow

import (
	"context"
	"encoding/json"
	"gin-idiary-appui/library/logit"
	followDao "gin-idiary-appui/modules/user/dao/follow"
	userModel "gin-idiary-appui/modules/user/model"
)

func Followings(ctx context.Context, pars userModel.FollowingPars) ([]int64, error) {
	var followingList []int64
	followings, err := followDao.GetFollowing(ctx, pars.UserID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(followings), &followingList)
	if err != nil {
		logit.Logger.Error("json.Unmarshal error: #+v", err)
		return nil, err
	}

	return followingList, nil
}
