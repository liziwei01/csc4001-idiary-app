/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 23:04:02
 * @Description: file content
 */
package services

import (
	"context"
	"encoding/json"
	diaryData "gin-idiary-appui/modules/diary/data"
	commentData "gin-idiary-appui/modules/diary/data/comment"
	diaryModel "gin-idiary-appui/modules/diary/model"
	uploadData "gin-idiary-appui/modules/upload/data"
	infoData "gin-idiary-appui/modules/user/data/info"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryListRequestPars) ([]diaryModel.DiaryInfoUnmarshaled, int64, error) {
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
	UnmarshaledDiary := make([]diaryModel.DiaryInfoUnmarshaled, 0)
	for k, v := range existedDiary {
		var newDiary diaryModel.DiaryInfoUnmarshaled
		newDiary.Address = v.Address
		newDiary.Authority = v.Authority
		newDiary.CommentCount = v.CommentCount
		newDiary.CommentList = v.CommentList
		newDiary.Content = v.Content
		newDiary.DBTime = v.DBTime
		newDiary.DeleteStatus = v.DeleteStatus
		newDiary.Device = v.Device
		newDiary.DiaryID = v.DiaryID
		newDiary.DislikeCount = v.DislikeCount
		newDiary.HasVoted = v.HasVoted
		newDiary.ImageList = make([]string, 0)
		newDiary.Nickname = v.Nickname
		newDiary.ReportCount = v.ReportCount
		newDiary.ShareCount = v.ShareCount
		newDiary.Tags = v.Tags
		newDiary.Title = v.Title
		newDiary.UserID = v.UserID
		newDiary.UserProfile = v.UserProfile
		newDiary.VoteCount = v.VoteCount

		if existedDiary[k].ImageList != "" {
			imgList := make([]string, 0)
			err = json.Unmarshal([]byte(existedDiary[k].ImageList), &imgList)
			if err != nil {
				return nil, 0, err
			}
			newDiary.ImageList = append(newDiary.ImageList, imgList...)
		}
		UnmarshaledDiary = append(UnmarshaledDiary, newDiary)
	}

	return UnmarshaledDiary, count, nil
}
