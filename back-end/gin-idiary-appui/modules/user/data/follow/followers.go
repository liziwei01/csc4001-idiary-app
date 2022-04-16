/*
 * @Author: liziwei01
 * @Date: 2022-04-16 18:24:49
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:16:36
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

func Followers(ctx context.Context, pars userModel.FollowerPars) ([]int64, error) {
	var followerList []int64
	followers, err := followDao.GetFollower(ctx, pars.UserID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(followers), &followerList)
	if err != nil {
		logit.Logger.Error("json.Unmarshal error: #+v", err)
		return nil, err
	}

	return followerList, nil
}
