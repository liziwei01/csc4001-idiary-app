/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-18 20:07:47
 * @Description: file content
 */
package follow

import (
	"context"
	followData "gin-idiary-appui/modules/user/data/follow"
	followModel "gin-idiary-appui/modules/user/model/follow"
)

func Followers(ctx context.Context, pars followModel.FollowerPars) ([]int64, error) {
	return followData.FollowersByEmail(ctx, pars.Email)
}
