/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-13 23:31:19
 * @Description: file content
 */
package data

import (
	"context"
	userDao "gin-idiary-appui/modules/user/dao"
	userModel "gin-idiary-appui/modules/user/model"
)

func CheckPassword(ctx context.Context, pars userModel.LoginPars) (bool, error) {
	pwd, err := userDao.GetPassword(ctx, pars)
	if err != nil {
		return false, err
	}
	return pwd == pars.Password, nil
}
