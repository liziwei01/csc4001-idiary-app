package model

type WorldInfo struct {
	UserID    string `ddb:"user_id" json:"user_id" `
	Title     string `ddb:"title" json:"title" `
	Content   string `ddb:"content" json:"content" `
	Timestamp int64  `ddb:"timestamp" json:"timestamp" `
	Authority string `ddb:"authority" json:"authority" `
	Address   string `ddb:"address" json:"address" `
}
