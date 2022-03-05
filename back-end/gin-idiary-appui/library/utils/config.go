/*
 * @Author: liziwei01
 * @Date: 2022-03-04 23:40:12
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-04 23:40:23
 * @Description: 解析配置
 */

package utils

import (
	"fmt"
	"gin-idiary-appui/library/conf"
	"os"
	"path/filepath"
)

// Get 解析配置.
func (u *UConfig) Get(confPath string, c interface{}) error {
	fileAbs, err := filepath.Abs(confPath)
	if err != nil {
		return err
	}
	if _, err := os.Stat(fileAbs); !os.IsNotExist(err) {
		if err := conf.Parse(fileAbs, c); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s conf file does not exist", confPath)
	}
	return nil
}
