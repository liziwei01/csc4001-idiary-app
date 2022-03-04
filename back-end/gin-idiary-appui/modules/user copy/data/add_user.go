/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:19:18
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-04 15:59:05
 * @Description: 数据处理
 */
package data

import (
	"context"
	userDao "gin-idiary-appui/modules/user/dao"
	userModel "gin-idiary-appui/modules/user/model"
)

func AddUser(ctx context.Context, pars userModel.UserRegisterPars) error {
	return userDao.AddUser(ctx, pars)
}
