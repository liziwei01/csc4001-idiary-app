/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:00:29
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 20:31:35
 * @Description: file content
 */
package data

import (
	"context"
	"gin-idiary-appui/library/redis"
	"gin-idiary-appui/modules/email/constant"
	"time"
)

const (
	EMAIL_VERIFICATION_CODE_EXPIRE_TIME = 60
)

// 存储验证码
func SaveVerificationCode(ctx context.Context, email string, code string) error {
	client, err := redis.GetClient(ctx, constant.SERVICE_CONF_REDIS_IDIARY)
	if err != nil {
		return err
	}

	key := getEmailKey(email)

	err = client.Set(ctx, key, code, getVerificationCodeExpireNanoSecond())
	if err != nil {
		return err
	}

	return nil
}

// 获取邮件验证码缓存过期时间
func getVerificationCodeExpireNanoSecond() time.Duration {
	return time.Second * time.Duration(EMAIL_VERIFICATION_CODE_EXPIRE_TIME)
}
