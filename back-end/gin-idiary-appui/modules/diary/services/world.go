package services

import (
	"context"
	diaryData "gin-idiary-appui/modules/diary/data"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func World(ctx context.Context) ([]diaryModel.DiaryInfo, error) {
	return diaryData.World(ctx)
}
