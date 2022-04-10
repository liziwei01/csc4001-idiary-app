package dao

import (
	"context"
	"gin-idiary-appui/library/mysql"
	userModel "gin-idiary-appui/modules/user/model"
)

func AddFriend(ctx context.Context, pars userModel.FriendsPars) error {
	client, err := mysql.GetClient(ctx, "idiary")
	if err != nil {
		return err
	}
	mapSliceInsertData := []map[string]interface{}{}
	mapSliceInsertData = append(mapSliceInsertData, map[string]interface{}{
		"user_id":   pars.UserID,
		"friend_id": pars.FriendID,

		// "nickname": pars.Nickname,
		// "email":    pars.Email,
	})
	_, err = client.Insert(ctx, "idiary_friends", mapSliceInsertData)
	if err != nil {
		return err
	}
	return nil
}
