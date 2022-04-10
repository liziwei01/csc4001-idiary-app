package data

import (
	"context"
	userDao "gin-idiary-appui/modules/user/dao"
	userModel "gin-idiary-appui/modules/user/model"
)

func Login(ctx context.Context, pars userModel.LoginPars) (bool, error) {
	return userDao.Login(ctx, pars)
}
