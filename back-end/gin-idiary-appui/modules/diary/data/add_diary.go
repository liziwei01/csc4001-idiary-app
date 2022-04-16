package data

import (
	"context"
	diaryDao "gin-idiary-appui/modules/diary/dao"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AddDiary(ctx context.Context, pars diaryModel.DiaryPars) error {
	return diaryDao.AddDiary(ctx, pars)
}
