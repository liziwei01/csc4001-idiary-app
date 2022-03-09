/*
 * @Author: liziwei01
 * @Date: 2022-03-04 23:38:24
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-09 21:31:17
 * @Description: file content
 */
package middleware

import (
	"context"
	"gin-idiary-appui/library/env"
	"gin-idiary-appui/library/utils"
	"path/filepath"

	rate "github.com/wallstreetcn/rate/redis"
)

const (
	middlewareConfPath  = "/middleware/"
	freqControlConfName = "freq_control.toml"
	signConfName        = "sign.toml"
	tokenConfName       = "token.toml"
)

var (
	freqControlConf *FreqControl
	tokenConf       *Token
	signConf        *Sign
	initConfig      = false

	getLimiter  *rate.Limiter
	postLimiter *rate.Limiter

	emailLimiter *rate.Limiter
)

type FreqControl struct {
	Enable bool

	GetLimit    int64
	PostLimit   int64
	PutLimit    int64
	DeleteLimit int64

	EmailLimit int64

	IpLimit int64
}
type Token struct {
	Enable      bool
	Token       string
	NoTokenPath []string
}

type Sign struct {
	Enable     bool
	Sign       string
	NoSignPath []string
}

func Init(ctx context.Context) {
	if initConfig == false {
		initConfig = true
		getConfig(freqControlConfName, &freqControlConf)
		getConfig(signConfName, &signConf)
		getConfig(tokenConfName, &tokenConf)
		// initFreqControl()
	}
}

func getConfig(confName string, conf interface{}) {
	confPath := filepath.Join(env.Default.ConfDir(), middlewareConfPath, confName)
	utils.Config.Get(confPath, conf)
}

func initFreqControl() {
	// initialize redis.
	rate.SetRedis(&rate.ConfigRedis{
		Host: "127.0.0.1",
		Port: 6379,
		Auth: "",
	})
}
