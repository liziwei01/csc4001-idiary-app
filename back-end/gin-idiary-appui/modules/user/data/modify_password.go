package data

import (
	"context"
	userDao "gin-idiary-appui/modules/user/dao"
	userModel "gin-idiary-appui/modules/user/model"
)

func ModifyPassword(ctx context.Context, pars userModel.UserPars) error {
	return userDao.ModifyPassword(ctx, pars)
}
