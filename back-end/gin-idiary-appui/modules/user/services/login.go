/*
 * @Author: liziwei01
 * @Date: 2022-04-18 17:27:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-18 20:49:40
 * @Description: file content
 */
package services

import (
	"context"
	infoData "gin-idiary-appui/modules/user/data/info"
	userModel "gin-idiary-appui/modules/user/model"
	infoModel "gin-idiary-appui/modules/user/model/info"
)

func Login(ctx context.Context, pars userModel.LoginPars) (infoModel.UserInfo, error) {
	return infoData.GetUserInfoByEmail(ctx, pars.Email)
}
