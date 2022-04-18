/*
 * @Author: liziwei01
 * @Date: 2022-04-18 20:15:40
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-18 21:49:13
 * @Description: file content
 */
package info

import (
	"context"
	uploadData "gin-idiary-appui/modules/upload/data"
	infoData "gin-idiary-appui/modules/user/data/info"
	infoModel "gin-idiary-appui/modules/user/model/info"
)

func Info(ctx context.Context, pars infoModel.UserInfoRequest) (infoModel.UserInfo, error) {
	info, err := infoData.GetUserInfoByEmail(ctx, pars.Email)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	info.Profile, err = uploadData.GetImageURL(ctx, info.Profile)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	return info, nil
}

func AllInfo(ctx context.Context, pars infoModel.UserInfoRequest) ([]infoModel.UserInfo, error) {
	infos, err := infoData.GetAllUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	for idx, info := range infos {
		infos[idx].Profile, err = uploadData.GetImageURL(ctx, info.Profile)
		if err != nil {
			return nil, err
		}
	}
	return infos, nil
}
