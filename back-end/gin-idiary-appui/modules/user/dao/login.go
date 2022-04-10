package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	userModel "gin-idiary-appui/modules/user/model"
)

func Login(ctx context.Context, pars userModel.LoginPars) (bool, error) {
	var password string

	// 数据库名字，之后替换
	client, err := mysql.GetClient(ctx, "idiary")

	where := map[string]interface{}{
		"user_id": pars.UserID,
	}

	columns := []string{"password"}

	// "idiary_diary"是表名，之后替换
	// diary是一个list
	_ = client.Query(ctx, "idiary_user", where, columns, password)

	if err != nil {
		return false, err
	}

	if password == pars.Password {
		return true, err
	} else {
		return false, err
	}

}
