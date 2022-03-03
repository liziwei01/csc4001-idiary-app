/*
 * @Author: liziwei01
 * @Date: 2022-03-03 19:55:32
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-03 20:01:12
 * @Description: 通过错误码表示错误
 */
package response

// 接口错误码.
const (
	// Success 成功.
	Success = 0
	// Failed 失败.
	Failed = -1
	// Unknown 未知错误.
	Unknown = -2
	// InvalidParams 参数错误.
	InvalidParams = -3
	// SignCheckFailed 接口数据签名校验失败.
	SignCheckFailed = -4
	// SignCheckFailed 接口数据token校验失败.
	TokenCheckFailed = -5
	// Insufficient authority.
	ERR_NO_AUTH = -6
	// Alert.
	Alert = -7
)

// 接口错误信息.
const (
	// msg: success.
	MsgSuccess = "Success"
	// msg: failed.
	MsgFailed = "Failure"
	// msg: unknown.
	MsgUnknown = "Unknown error"
	// msg: invalid params.
	MsgInvalidParams = "Invalid params"
	// msg: invalid md5.
	MsgSignCheckFailed = "Sign check failed"
	// msg: invalid token.
	MsgTokenCheckFailed = "Token check failed"
	// msg: user account closed.
	MsgUserAccountClosed = "User account closed"
)

// map: code -> msg.
var CodeMsgMap = map[int]string{
	Success:          MsgSuccess,
	Failed:           MsgFailed,
	Unknown:          MsgUnknown,
	InvalidParams:    MsgInvalidParams,
	SignCheckFailed:  MsgSignCheckFailed,
	TokenCheckFailed: MsgTokenCheckFailed,

	// Authority.
	ERR_NO_AUTH: "权限不足: 请hi超级管理员操作或联系本配置的创建人",

	// Alert.
	Alert: "此接口仅供监控、配置Review等离线流量使用, 严禁接入线上流量！",
}
