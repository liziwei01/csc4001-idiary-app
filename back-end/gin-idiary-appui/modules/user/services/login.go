package services

import (
	"context"
	userData "gin-idiary-appui/modules/user/data"
	userModel "gin-idiary-appui/modules/user/model"
)

func Login(ctx context.Context, pars userModel.LoginPars) (bool, error) {
	return userData.Login(ctx, pars)
}
