package data

import (
	"context"
	diaryDao "csc4001-idiary-app/back-end/gin-idiary-appui/modules/diary/dao"
	diaryModel "csc4001-idiary-app/back-end/gin-idiary-appui/modules/diary/model"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryRegisterPars) error {
	return diaryDao.AllDiary(ctx, pars)
}
