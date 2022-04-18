package services

import (
	"context"
	userData "gin-idiary-appui/modules/user/data"
	userModel "gin-idiary-appui/modules/user/model"
)

func ModifyPassword(ctx context.Context, pars userModel.UserPars) error {
	return userData.ModifyPassword(ctx, pars)
}
