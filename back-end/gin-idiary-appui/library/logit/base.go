/*
 * @Author: liziwei01
 * @Date: 2022-03-04 15:40:52
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-04 15:40:53
 * @Description: file content
 */
package logit

import (
	lib "github.com/baidu/go-lib/log"
)

var (
	Logger = &lib.Logger
)

/**
 * @description: all the log are recorded under ./log
 * @param {string} programName
 * @return {*}
 */
func Init(programName string) error {
	return initLog(programName)
}

/**
 * @description:
 * @param {string} programName
 * @return {*}
 */
func initLog(programName string) error {
	err := lib.Init(programName, "INFO", "./log", true, "H", 5)
	if err != nil {
		return err
	}
	return nil
}
