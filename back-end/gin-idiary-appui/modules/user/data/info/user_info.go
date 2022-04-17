/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:12:20
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:25:05
 * @Description: file content
 */
/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:12:20
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:12:20
 * @Description: file content
 */
package info

import (
	"context"
	"fmt"
	infoDao "gin-idiary-appui/modules/user/dao/info"
	infoModel "gin-idiary-appui/modules/user/model/info"
)

func BatchGetUserInfo(ctx context.Context, userIDs []int64) ([]infoModel.UserInfo, error) {
	userInfos := make([]infoModel.UserInfo, 0)

	for _, userID := range userIDs {
		if userID <= 0 {
			return nil, fmt.Errorf("invalid user id: %d", userID)
		}
		userInfo, err := infoDao.GetUserInfoByUserID(ctx, userID)
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}
