/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 11:02:15
 * @Description: file content
 */
package data

import (
	"context"
	userDao "gin-idiary-appui/modules/user/dao"
	userModel "gin-idiary-appui/modules/user/model"
)

func ModifyPassword(ctx context.Context, pars userModel.UserPars) error {
	return userDao.ModifyPassword(ctx, pars)
}
