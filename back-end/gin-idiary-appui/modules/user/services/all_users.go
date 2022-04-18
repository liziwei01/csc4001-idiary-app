package services

import (
	"context"
	userData "gin-idiary-appui/modules/user/data"
	userModel "gin-idiary-appui/modules/user/model"
)

func AllUsers(ctx context.Context, pars userModel.UserListRequestPars) ([]userModel.UserInfo, int64, error) {

	users, count, err := userData.AllUsers(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	// existedUser := make([]userModel.UserInfo, 0)
	// // 3 get user info
	// userIDs := make([]int64, 0)
	// for _, v := range existedUser {
	// 	userIDs = append(userIDs, v.UserID)
	// }

	// userInfos, err := infoData.BatchGetUserInfo(ctx, userIDs)
	// if err != nil {
	// 	return nil, 0, err
	// }

	// // 6 get profile url
	// for k, v := range existedUser {
	// 	existedUser[k].Profile, err = uploadData.GetImageURL(ctx, v.Profile)
	// 	if err != nil {
	// 		return nil, 0, err
	// 	}

	// }

	// // 7 unmarshal image_list
	// UnmarshaledUser := make([]userModel.UserInfo, 0)
	// for v := range existedUser {
	// 	var newUser userModel.UserInfo
	// 	newUser.Address = v.Address
	// 	newUser.Email = v.Email
	// 	newUser.Password = v.Password
	// 	newUser.Nickname = v.CommentList
	// 	newUser.Profile = v.Profile
	// 	newUser.DBTime = v.DBTime

	// 	UnmarshaledUser = append(UnmarshaledIser, newUser)
	// }

	return users, count, nil
}
