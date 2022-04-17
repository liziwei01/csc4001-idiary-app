/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:42:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:44:43
 * @Description: file content
 */
package comment

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/diary/constant"
	commentModel "gin-idiary-appui/modules/diary/model/comment"
)

const (
	DIARY_COMMENT_TABLE = "tb_diary_comment"
)

func GetCommentByDiaryID(ctx context.Context, diaryID int64) ([]commentModel.CommentInfo, error) {
	var diary []commentModel.CommentInfo

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_FEED)
	if err != nil {
		return nil, err
	}

	tableName := DIARY_COMMENT_TABLE

	where := map[string]interface{}{
		"_orderby": "db_time desc",
		"diary_id": diaryID,
	}

	columns := []string{"*"}

	err = client.Query(ctx, tableName, where, columns, &diary)
	if err != nil {
		return nil, err
	}

	return diary, nil
}

func GetCommentByDiaryIDCount(ctx context.Context, diaryID int64) (int64, error) {
	var count = make([]struct {
		Count int64 `ddb:"count"`
	}, 1)

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_FEED)
	if err != nil {
		return 0, err
	}

	tableName := DIARY_COMMENT_TABLE

	where := map[string]interface{}{
		"_orderby": "db_time desc",
		"diary_id": diaryID,
	}

	columns := []string{"count(1) as count"}

	err = client.Query(ctx, tableName, where, columns, &count)
	if err != nil {
		return 0, err
	}

	return count[0].Count, nil
}
