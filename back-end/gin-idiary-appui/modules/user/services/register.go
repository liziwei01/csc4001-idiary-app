package services

import (
	"context"
	userData "gin-idiary-appui/modules/user/data"
	userModel "gin-idiary-appui/modules/user/model"
)

func Register(ctx context.Context, pars userModel.RegisterPars) error {
	return userData.Register(ctx, pars)
}
