/*
 * @Author: liziwei01
 * @Date: 2022-03-04 23:38:24
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-05 00:24:47
 * @Description: file content
 */
package middleware

import (
	"context"
	"gin-idiary-appui/library/env"
	"gin-idiary-appui/library/utils"
	"path/filepath"
)

const (
	middlewareConfPath = "/middleware/"
	filterConfName     = "filter.toml"
)

var (
	config     *Config
	initConfig = false
)

type Config struct {
	FreqControl struct {
		Enable bool
	}
	// token 的配置.
	Token struct {
		Token       string
		NoTokenPath []string
		Sign        string
		NoSignPath  []string
	}
	Sign struct {
		Sign       string
		NoSignPath []string
	}
}

func Init(ctx context.Context) {
	if initConfig == false {
		initConfig = true
		config = GetFilterConfig()
	}
}

func GetFilterConfig() *Config {
	var c *Config
	confPath := filepath.Join(env.Default.ConfDir(), middlewareConfPath, filterConfName)
	utils.Config.Get(confPath, &c)
	return c
}
