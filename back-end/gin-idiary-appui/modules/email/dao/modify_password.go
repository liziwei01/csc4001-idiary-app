package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	emailModel "gin-idiary-appui/modules/email/model"
)

func ModifyPassword(ctx context.Context, pars emailModel.UserPars, password string) error {

	// 数据库名字，之后替换
	client, err := mysql.GetClient(ctx, "idiary")

	where := map[string]interface{}{
		"user_id": pars.UserID,
	}

	update_password := map[string]interface{}{
		"password": password,
	}

	// "idiary_diary"是表名，之后替换
	// diary是一个list
	_, err = client.Update(ctx, "idiary_user", where, update_password)
	if err != nil {
		return err
	}
	return err
}
