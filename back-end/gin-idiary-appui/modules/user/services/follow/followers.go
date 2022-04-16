/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:16:48
 * @Description: file content
 */
package follow

import (
	"context"
	followData "gin-idiary-appui/modules/user/data/follow"
	userModel "gin-idiary-appui/modules/user/model"
)

func Followers(ctx context.Context, pars userModel.FollowerPars) ([]int64, error) {
	return followData.Followers(ctx, pars)
}
