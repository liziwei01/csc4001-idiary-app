/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:16:06
 * @Description: file content
 */
package follow

import (
	"context"
	followData "gin-idiary-appui/modules/user/data/follow"
	followModel "gin-idiary-appui/modules/user/model/follow"
)

func Follow(ctx context.Context, pars followModel.FollowPars) error {
	return followData.Follow(ctx, pars.UserID, pars.FollowingID)
}
