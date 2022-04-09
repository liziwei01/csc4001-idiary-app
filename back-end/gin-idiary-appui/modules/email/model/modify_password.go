package model

type UserPars struct {
	UserID      string `form:"user_id" json:"user_id" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	NewPassword string `form:"newPassword" json:"newPassword" binding:"required"`
	Nickname    string `form:"nickname" json:"nickname" binding:"required"`
	City        string `form:"city" json:"city" binding:"required"`
	Picture     string `form:"picture" json:"picture" binding:"required"`
}

type UserInfo struct {
	UserID   string `ddb:"user_id" json:"user_id" `
	Password string `ddb:"password" json:"password" `
	Nickname string `ddb:"nickname" json:"nickname" `
	City     string `ddb:"city" json:"city" `
	Picture  string `ddb:"picture" json:"picture" `
}
