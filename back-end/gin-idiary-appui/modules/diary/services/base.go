/*
 * @Author: liziwei01
 * @Date: 2022-04-22 20:22:36
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-25 21:18:14
 * @Description: file content
 */
package services

import (
	"context"
	"encoding/json"
	"gin-idiary-appui/library/utils"
	commentData "gin-idiary-appui/modules/diary/data/comment"
	diaryModel "gin-idiary-appui/modules/diary/model"
	uploadData "gin-idiary-appui/modules/upload/data"
	followData "gin-idiary-appui/modules/user/data/follow"
	infoData "gin-idiary-appui/modules/user/data/info"
)

func ProcessDiary(ctx context.Context, requestUID int64, diaries []diaryModel.DiaryInfo, private ...bool) ([]diaryModel.DiaryInfoUnmarshaled, error) {
	publicOnly := false
	if len(private) > 0 {
		publicOnly = private[0]
	}

	// process 1: hide diary (diaries -> existDiary)
	existedDiary := make([]diaryModel.DiaryInfo, 0)
	for _, v := range diaries {
		// (1) hide deleted diary
		if v.DeleteStatus == true {
			continue
		}
		// (2) hide private and protected diary
		if publicOnly && v.Authority != 0 {
			continue
		}
		existedDiary = append(existedDiary, v)
	}

	// process 2: append raw comments
	// (1) get comments
	for k, v := range existedDiary {
		comments, count, err := commentData.GetCommentByDiaryID(ctx, v.DiaryID)
		if err != nil {
			return nil, err
		}
		existedDiary[k].CommentList = comments
		existedDiary[k].CommentCount = count
	}

	// process 3: append like status
	// (1) get likes status
	for k, v := range existedDiary {
		likeCount, err := commentData.DiaryLikeCount(ctx, v.DiaryID)
		if err != nil {
			return nil, err
		}
		hasLiked, err := commentData.DiaryHasLiked(ctx, requestUID, v.DiaryID)
		if err != nil {
			return nil, err
		}
		existedDiary[k].LikeCount = likeCount
		existedDiary[k].HasLiked = hasLiked
	}

	// process 4: show user info
	// (1) get diary user info
	 userIDs := make([]int64, 0)
	 for _, v := range existedDiary {
	 	userIDs = append(userIDs, v.UserID)
	 }
	 userInfos, err := infoData.BatchGetUserInfo(ctx, userIDs)
	 if err != nil {
	 	return nil, err
	 }

	// (2) combine diary user info and diary
	 for k := range existedDiary {
	 	existedDiary[k].Nickname = userInfos[k].Nickname
	 	existedDiary[k].UserProfile = userInfos[k].Profile
	 }

	// (3) get comment user info and combine
	for k, v := range existedDiary {
		comments := v.CommentList
		for idx, comment := range comments {
			userInfo, err := infoData.GetUserInfoByUserID(ctx, comment.UserID)
			if err != nil {
				return nil, err
			}
			comments[idx].Nickname = userInfo.Nickname
			comments[idx].Profile = userInfo.Profile
		}
		existedDiary[k].CommentList = comments
	}

	// process 5: give a temporary url to all images (existDiary -> diariesUnmarshaled)
	// (1) get user profile url
	for k, v := range existedDiary {
		existedDiary[k].UserProfile, err = uploadData.GetImageURL(ctx, v.UserProfile)
		if err != nil {
			return nil, err
		}

		for _, comment := range v.CommentList {
			comment.Profile, err = uploadData.GetImageURL(ctx, comment.Profile)
			if err != nil {
				return nil, err
			}
		}
	}

	// (2) unmarshal diary image_list and get url
	diariesUnmarshaled := make([]diaryModel.DiaryInfoUnmarshaled, 0)
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
		newDiary.HasVoted = v.HasLiked
		newDiary.ImageList = make([]string, 0)
		newDiary.Nickname = v.Nickname
		newDiary.ReportCount = v.ReportCount
		newDiary.ShareCount = v.ShareCount
		newDiary.Tags = v.Tags
		newDiary.Title = v.Title
		newDiary.UserID = v.UserID
		newDiary.UserProfile = v.UserProfile
		newDiary.VoteCount = v.LikeCount

		if existedDiary[k].ImageList != "" && existedDiary[k].ImageList != "[]" && existedDiary[k].ImageList != "\"\"" {
			imgList := make([]string, 0)
			err = json.Unmarshal([]byte(existedDiary[k].ImageList), &imgList)
			if err != nil {
				return nil, err
			}
			for idx, imageName := range imgList {
				imgList[idx], err = uploadData.GetImageURL(ctx, imageName)
				if err != nil {
					return nil, err
				}
			}

			newDiary.ImageList = append(newDiary.ImageList, imgList...)
		}
		diariesUnmarshaled = append(diariesUnmarshaled, newDiary)
	}

	// process 6: user follow relationship
	// (1) get followings
	followings, err := followData.Followings(ctx, requestUID)
	if err != nil {
		return nil, err
	}

	// (2) calculate user following relationship
	for k, v := range diariesUnmarshaled {
		diariesUnmarshaled[k].HasFollowed = utils.Slice.In(v.UserID, followings)
		if err != nil {
			return nil, err
		}
	}
	return diariesUnmarshaled, nil
}
