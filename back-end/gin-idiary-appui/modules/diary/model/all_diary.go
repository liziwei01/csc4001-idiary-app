/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 18:50:52
 * @Description: file content
 */
package model

type DiaryListRequestPars struct {
	UserID     string `form:"user_id" json:"user_id"`
	PageIndex  uint   `form:"page_index" json:"page_index"`
	PageLength uint   `form:"page_length" json:"page_length"`
}

type DiaryInfo struct {
	UserID    string `ddb:"user_id" json:"user_id" `
	Title     string `ddb:"title" json:"title" `
	Content   string `ddb:"content" json:"content" `
	Timestamp int64  `ddb:"timestamp" json:"timestamp" `
	Authority string `ddb:"authority" json:"authority" `
	Address   string `ddb:"address" json:"address" `
}
