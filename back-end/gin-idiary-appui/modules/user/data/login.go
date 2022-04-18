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
