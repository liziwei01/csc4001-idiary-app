/*
 * @Author: liziwei01
 * @Date: 2022-04-24 20:37:32
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-25 09:53:58
 * @Description: file content
 */
package comment

import (
	"context"
	"fmt"
	"gin-idiary-appui/library/logit"
	commentData "gin-idiary-appui/modules/diary/data/comment"
	commentModel "gin-idiary-appui/modules/diary/model/comment"
)

func Like(ctx context.Context, pars commentModel.DiaryLikeRequestPars) error {
	// 检查 redis 是否存在 key: 即是否已经like过
	hasLiked, err := commentData.DiaryHasLiked(ctx, pars.UserID, pars.DiaryID)
	if err != nil {
		return err
	}
	// 1 点赞
	if !pars.ShouldUnlike {
		// 如果存在, 则报错
		if hasLiked {
			err = fmt.Errorf("已经点赞过")
			logit.Logger.Info(err.Error())
			return err
		}

		// 如果不存在, 点赞
		err = commentData.Like(ctx, pars.UserID, pars.DiaryID)
		if err != nil {
			return err
		}

		// 2 取消点赞
	} else {
		if !hasLiked {
			err = fmt.Errorf("已经取消点赞")
			logit.Logger.Info(err.Error())
			return err
		}

		// 如果存在, 取消点赞
		err = commentData.UnLike(ctx, pars.UserID, pars.DiaryID)
		if err != nil {
			return err
		}
	}

	return nil
}
