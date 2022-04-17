/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 18:02:14
 * @Description: file content
 */
package services

import (
	"context"
	diaryData "gin-idiary-appui/modules/diary/data"
	commentData "gin-idiary-appui/modules/diary/data/comment"
	diaryModel "gin-idiary-appui/modules/diary/model"
	uploadData "gin-idiary-appui/modules/upload/data"
	infoData "gin-idiary-appui/modules/user/data/info"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryListRequestPars) ([]diaryModel.DiaryInfo, int64, error) {
	// 1 get diary
	diaries, count, err := diaryData.AllDiary(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	// 2 remove deleted diary
	existedDiary := make([]diaryModel.DiaryInfo, 0)
	for _, v := range diaries {
		if v.DeleteStatus == true {
			continue
		}
		existedDiary = append(existedDiary, v)
	}

	// 3 get user info
	userIDs := make([]int64, 0)
	for _, v := range existedDiary {
		userIDs = append(userIDs, v.UserID)
	}

	userInfos, err := infoData.BatchGetUserInfo(ctx, userIDs)
	if err != nil {
		return nil, 0, err
	}

	// 4 combine user info and diary
	for k := range existedDiary {
		existedDiary[k].Nickname = userInfos[k].Nickname
		existedDiary[k].UserProfile = userInfos[k].Profile
	}

	// 5 get comments
	for k, v := range existedDiary {
		comments, count, err := commentData.GetCommentWithUserInfoByDiaryID(ctx, v.UserID)
		if err != nil {
			return nil, 0, err
		}
		existedDiary[k].CommentList = comments
		existedDiary[k].CommentCount = count
	}

	// 6 get profile url
	for k, v := range existedDiary {
		existedDiary[k].UserProfile, err = uploadData.GetImageURL(ctx, v.UserProfile)
		if err != nil {
			return nil, 0, err
		}
	}

	// 7 unmarshal image_list
	// for k := range existedDiary {
	// 	if existedDiary[k].ImageList != "" {
	// 		imgList := make([]diaryModel.DiaryInfo, 0)
	// 		err = json.Unmarshal([]byte(existedDiary[k].ImageList), &imgList)
	// 		if err != nil {
	// 			return nil, 0, err
	// 		}
	// 		existedDiary[k].ImageList = imgList
	// 	}
	// }

	return existedDiary, count, nil
}
