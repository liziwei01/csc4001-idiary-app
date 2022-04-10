package model

type LoginPars struct {
	UserID   string `form:"user_id" json:"user_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserInfo struct {
	UserID   string `ddb:"user_id" json:"user_id" `
	Password string `ddb:"password" json:"password" `
}
