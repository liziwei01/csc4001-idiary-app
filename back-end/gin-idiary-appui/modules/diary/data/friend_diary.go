/*
 * @Author: liziwei01
 * @Date: 2022-04-16 20:06:15
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 17:09:19
 * @Description: file content
 */

package data

import (
	"context"
	diaryDao "gin-idiary-appui/modules/diary/dao"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func FriendDiary(ctx context.Context, pars diaryModel.FriendDiaryListRequestPars, friendIDs []int64) ([]diaryModel.DiaryInfo, int64, error) {
	diaries, err := diaryDao.BatchGetDiary(ctx, pars, friendIDs)
	if err != nil {
		return nil, 0, err
	}

	count, err := diaryDao.BatchGetDiaryCount(ctx, pars, friendIDs)
	if err != nil {
		return nil, 0, err
	}

	return diaries, count, err
}

func CalcFriend(ctx context.Context, followingList, followerList []int64) ([]int64, error) {
	friendList := make([]int64, 0)
	for _, follow := range followingList {
		for _, fans := range followerList {
			if follow == fans {
				friendList = append(friendList, follow)
			}
		}
	}
	return friendList, nil
}
