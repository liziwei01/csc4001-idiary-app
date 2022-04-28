/*
 * @Author: liziwei01
 * @Date: 2022-04-09 23:52:44
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 19:35:35
 * @Description: file content
 */
package services

import (
	"context"
	diaryData "gin-idiary-appui/modules/diary/data"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func DeleteDiary(ctx context.Context, pars diaryModel.DeleteDiaryRequestPars) error {
	return diaryData.DeleteDiary(ctx, pars.DiaryID)
}
