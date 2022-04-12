package model

type RegisterPars struct {
	UserID   string `form:"user_id" json:"user_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	City     string `form:"city" json:"city" binding:"required"`
	Picture  string `form:"picture" json:"picture" binding:"required"`
}
