/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:46
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-03 16:22:41
 * @Description: file content
 */

package httpapi

import (
	"io"
	"net/http"

	userRouters "gin-idiary-appui/modules/user/routers"
)

/**
 * @description: start http server and start listening
 * @param {*}
 * @return {*}
 */
func InitRouters() {
	// init routers
	userRouters.Init()

	// safe router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello! THis is iDiary. Welcome to our website!")
	})
}
