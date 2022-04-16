/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:16:23
 * @Description: file content
 */
package follow

import (
	"context"
	followData "gin-idiary-appui/modules/user/data/follow"
	userModel "gin-idiary-appui/modules/user/model"
)

func Followings(ctx context.Context, pars userModel.FollowingPars) ([]int64, error) {
	return followData.Followings(ctx, pars)
}
