package model

type FriendsPars struct {
	UserID   string `form:"user_id" json:"user_id" binding:"required"`
	FriendID string `form:"friend_id" json:"friend_id" binding:"required"`
}
