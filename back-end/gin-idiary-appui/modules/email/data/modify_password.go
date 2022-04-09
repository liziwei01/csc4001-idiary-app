package data

import (
	"context"
	emailDao "gin-idiary-appui/modules/email/dao"
	emailModel "gin-idiary-appui/modules/email/model"
)

func ModifyPassword(ctx context.Context, pars emailModel.UserPars) error {
	return emailDao.ModifyPassword(ctx, pars)
}
