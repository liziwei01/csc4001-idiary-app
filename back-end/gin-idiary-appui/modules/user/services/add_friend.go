package services

import (
	"context"
	userData "gin-idiary-appui/modules/user/data"
	userModel "gin-idiary-appui/modules/user/model"
)

func AddFriend(ctx context.Context, pars userModel.FriendsPars) error {
	return userData.AddFriend(ctx, pars)
}
