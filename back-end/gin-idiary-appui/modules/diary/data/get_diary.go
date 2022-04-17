/*
 * @Author: liziwei01
 * @Date: 2022-04-17 16:54:13
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 17:07:19
 * @Description: file content
 */
package data

import (
	"context"
	diaryDao "gin-idiary-appui/modules/diary/dao"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func GetDiary(ctx context.Context, pars diaryModel.DiaryListRequestPars) ([]diaryModel.DiaryInfo, int64, error) {
	diaries, err := diaryDao.GetDiaryByUserID(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	count, err := diaryDao.GetDiaryByUserIDCount(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return diaries, count, err
}
