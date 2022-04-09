package model

type DiaryShowPars struct {
	UserID    string `form:"user_id" json:"user_id" binding:"required" `
	Title     string `form:"title" json:"title" `
	Content   string `form:"content" json:"content" `
	Timestamp int64  `form:"timestamp" json:"timestamp" `
	Authority string `form:"authority" json:"authority" `
	Address   string `form:"address" json:"address" `
}

type DiaryInfo struct {
	UserID    string `ddb:"user_id" json:"user_id" `
	Title     string `ddb:"title" json:"title" `
	Content   string `ddb:"content" json:"content" `
	Timestamp int64  `ddb:"timestamp" json:"timestamp" `
	Authority string `ddb:"authority" json:"authority" `
	Address   string `ddb:"address" json:"address" `
}
