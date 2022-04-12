/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 15:26:22
 * @Description: file content
 */
package services

import (
	"context"
	uploadData "gin-idiary-appui/modules/upload/data"
	uploadModel "gin-idiary-appui/modules/upload/model"
)

// 图片上传
func UploadImage(ctx context.Context, pars uploadModel.UploadPars) (string, error) {
	// 上传oss
	src, err := uploadData.UploadOSS(ctx, pars.UserID, pars.FileHeader)
	if err != nil {
		return "", err
	}

	return src, nil
}
