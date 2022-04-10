package data

import (
	"context"
	userDao "gin-idiary-appui/modules/user/dao"
	userModel "gin-idiary-appui/modules/user/model"
)

func AddFriend(ctx context.Context, pars userModel.FriendsPars) error {
	return userDao.AddFriend(ctx, pars)
}
