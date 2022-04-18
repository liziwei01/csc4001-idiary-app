/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-18 17:31:18
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
	followData "gin-idiary-appui/modules/user/data/follow"
	infoData "gin-idiary-appui/modules/user/data/info"
)

func FriendDiary(ctx context.Context, pars diaryModel.FriendDiaryListRequestPars) ([]diaryModel.DiaryInfoUnmarshaled, int64, error) {
	// 1 获取关注列表
	followingList, err := followData.Followings(ctx, pars.UserID)
	if err != nil {
		return nil, 0, err
	}

	// 2 获取粉丝列表
	followerList, err := followData.Followers(ctx, pars.UserID)
	if err != nil {
		return nil, 0, err
	}

	// 3 计算 朋友：关注列表 与 粉丝列表重合的部分
	if len(followerList) == 0 || len(followingList) == 0 {
		return nil, 0, nil
	}

	friendIDs, err := diaryData.CalcFriend(ctx, followingList, followerList)
	if err != nil {
		return nil, 0, err
	}

	// 4 获取朋友的日记列表
	diaryList, total, err := diaryData.FriendDiary(ctx, pars, friendIDs)
	if err != nil {
		return nil, 0, err
	}

	// 5 remove deleted diary
	existedDiary := make([]diaryModel.DiaryInfo, 0)
	for _, v := range diaryList {
		if v.DeleteStatus == true {
			continue
		}
		existedDiary = append(existedDiary, v)
	}

	// 6 get user info
	userIDs := make([]int64, 0)
	for _, v := range existedDiary {
		userIDs = append(userIDs, v.UserID)
	}

	userInfos, err := infoData.BatchGetUserInfo(ctx, userIDs)
	if err != nil {
		return nil, 0, err
	}

	// 7 combine user info and diary
	for k, v := range existedDiary {
		v.Nickname = userInfos[k].Nickname
		v.UserProfile = userInfos[k].Profile
	}

	// 8 get comments
	for k, v := range existedDiary {
		comments, count, err := commentData.GetCommentByDiaryID(ctx, v.UserID)
		if err != nil {
			return nil, 0, err
		}
		existedDiary[k].CommentList = comments
		existedDiary[k].CommentCount = count
	}

	// 9 get profile url
	for k, v := range existedDiary {
		existedDiary[k].UserProfile, err = uploadData.GetImageURL(ctx, v.UserProfile)
		if err != nil {
			return nil, 0, err
		}

		for _, comment := range v.CommentList {
			comment.Profile, err = uploadData.GetImageURL(ctx, comment.Profile)
			if err != nil {
				return nil, 0, err
			}
		}
	}

	// 10 unmarshal image_list
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

			for idx, imageName := range imgList {
				imgList[idx], err = uploadData.GetImageURL(ctx, imageName)
				if err != nil {
					return nil, 0, err
				}
			}

			newDiary.ImageList = append(newDiary.ImageList, imgList...)
		}
		UnmarshaledDiary = append(UnmarshaledDiary, newDiary)
	}

	return UnmarshaledDiary, total, nil
}
