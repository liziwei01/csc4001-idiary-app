package services

import (
	"context"
	emailData "gin-idiary-appui/modules/email/data"
	emailModel "gin-idiary-appui/modules/email/model"
)

func ModifyPassword(ctx context.Context, pars emailModel.UserPars, password string) error {
	return emailData.ModifyPassword(ctx, pars, password)
}
