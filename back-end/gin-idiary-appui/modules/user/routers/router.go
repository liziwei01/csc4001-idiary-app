/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:16:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-03 16:23:00
 * @Description: file content
 */

package routers

import (
	"net/http"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init() {
	http.HandleFunc("/addUser", func(rw http.ResponseWriter, r *http.Request) {})
}
