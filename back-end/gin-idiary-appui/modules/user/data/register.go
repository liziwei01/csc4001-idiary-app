package data

import (
	"context"
	userDao "gin-idiary-appui/modules/user/dao"
	userModel "gin-idiary-appui/modules/user/model"
)

func Register(ctx context.Context, pars userModel.RegisterPars) error {
	return userDao.Register(ctx, pars)
}
