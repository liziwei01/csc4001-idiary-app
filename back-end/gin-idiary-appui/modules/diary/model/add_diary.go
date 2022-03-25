package model

type DiaryRegisterPars struct {
	UserID    string `form:"user_id" json:"user_id" binding:"required"`
	Title     string `form:"title" json:"title" binding:"required"`
	Content   string `form:"content" json:"content" binding:"required"`
	Timestamp int64  `form:"timestamp" json:"timestamp" binding:"required"`
	Authority string `form:"authority" json:"authority" binding:"required"`
	Address   string `form:"address" json:"address" binding:"required"`
}
