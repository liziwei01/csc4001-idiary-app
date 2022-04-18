package data

import (
	"context"
	userDao "gin-idiary-appui/modules/user/dao"
	userModel "gin-idiary-appui/modules/user/model"
)

func AllUsers(ctx context.Context, pars userModel.UserListRequestPars) ([]userModel.UserInfo, int64, error) {
	users, err := userDao.AllUsers(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	count, err := userDao.AllUserCount(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return users, count, err
}
