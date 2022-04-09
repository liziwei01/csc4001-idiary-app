package data

import (
	"context"
	diaryDao "gin-idiary-appui/modules/diary/dao"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryRegisterPars) ([]diaryModel.DiaryInfo, error) {
	return diaryDao.AllDiary(ctx, pars)
}
