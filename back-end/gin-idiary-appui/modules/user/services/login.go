/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-13 23:31:37
 * @Description: file content
 */
package services

import (
	"context"
	userData "gin-idiary-appui/modules/user/data"
	userModel "gin-idiary-appui/modules/user/model"
)

func Login(ctx context.Context, pars userModel.LoginPars) (bool, error) {
	return userData.CheckPassword(ctx, pars)
}
