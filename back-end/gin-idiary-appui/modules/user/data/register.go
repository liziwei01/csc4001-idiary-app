/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:21:06
 * @Description: file content
 */
package data

import (
	"context"
	userDao "gin-idiary-appui/modules/user/dao"
	infoDao "gin-idiary-appui/modules/user/dao/info"
	userModel "gin-idiary-appui/modules/user/model"
)

func HasRegisted(ctx context.Context, email string) (bool, error) {
	userInfos, err := infoDao.GetUserInfoByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return len(userInfos) > 0, nil
}

func Register(ctx context.Context, pars userModel.RegisterPars) error {
	return userDao.Register(ctx, pars)
}
