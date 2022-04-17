/*
 * @Author: liziwei01
 * @Date: 2022-04-17 15:55:22
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 16:01:24
 * @Description: file content
 */
package services

import (
	"context"
	infoData "gin-idiary-appui/modules/user/data/info"

	userModel "gin-idiary-appui/modules/user/model"
)

func ModifyProfile(ctx context.Context, pars userModel.Profile) (err error) {
	return infoData.ModifyProfile(ctx, pars.UserID, pars.UserProfile)
}
