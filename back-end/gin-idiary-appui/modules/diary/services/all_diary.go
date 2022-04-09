package services

import (
	"context"
	diaryData "gin-idiary-appui/modules/diary/data"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryShowPars) ([]diaryModel.DiaryInfo, error) {
	return diaryData.AllDiary(ctx, pars)
}
