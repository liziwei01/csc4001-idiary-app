/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:20:19
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-04 16:00:10
 * @Description: 完整服务
 */
package services

import (
	"context"
	userData "gin-idiary-appui/modules/user/data"
	userModel "gin-idiary-appui/modules/user/model"
)

func AddUser(ctx context.Context, pars userModel.UserRegisterPars) error {
	return userData.AddUser(ctx, pars)
}
