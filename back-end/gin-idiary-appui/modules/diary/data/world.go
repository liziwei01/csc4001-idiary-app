package data

import (
	"context"
	diaryDao "gin-idiary-appui/modules/diary/dao"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func World(ctx context.Context) ([]diaryModel.DiaryInfo, error) {
	return diaryDao.World(ctx)
}
