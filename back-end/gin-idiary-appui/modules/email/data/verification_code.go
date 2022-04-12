/*
 * @Author: liziwei01
 * @Date: 2022-04-12 13:51:55
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 13:55:27
 * @Description: file content
 */
package data

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gogf/gf/util/gconv"
)

func GetRandomVerificationCode(ctx context.Context, email string) (string, error) {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano()+gconv.Int64(email))).Int31n(1000000))
	return code, nil
}
