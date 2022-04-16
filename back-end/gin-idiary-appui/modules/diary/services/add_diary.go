package services

import (
	"context"
	diaryData "gin-idiary-appui/modules/diary/data"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AddDiary(ctx context.Context, pars diaryModel.DiaryPars) error {
	return diaryData.AddDiary(ctx, pars)
}
