/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 11:01:46
 * @Description: file content
 */
package services

import (
	"context"
	userData "gin-idiary-appui/modules/user/data"
	userModel "gin-idiary-appui/modules/user/model"
)

func ModifyPassword(ctx context.Context, pars userModel.UserPars) error {
	return userData.ModifyPassword(ctx, pars)
}
