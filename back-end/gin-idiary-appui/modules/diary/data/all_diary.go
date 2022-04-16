/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:50:45
 * @Description: file content
 */
package data

import (
	"context"
	diaryDao "gin-idiary-appui/modules/diary/dao"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryListRequestPars) ([]diaryModel.DiaryInfo, int64, error) {
	diaries, err := diaryDao.AllDiary(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	count, err := diaryDao.AllDiaryCount(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return diaries, count, err
}
