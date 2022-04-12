/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 14:22:28
 * @Description: file content
 */
package services

import (
	"context"
	"fmt"
	emailData "gin-idiary-appui/modules/email/data"
	emailModel "gin-idiary-appui/modules/email/model"
)

func VerificationCode(ctx context.Context, pars emailModel.EmailPars) error {
	// 检查 redis 是否存在 key
	exists, err := emailData.EmailSent(ctx, pars.Email)
	if err != nil {
		return err
	}

	// 如果存在, 则报错
	if exists {
		return fmt.Errorf("该邮箱60s内已发送过验证码")
	}

	// 如果不存在, 发送验证码
	// 获取随机六位验证码
	code, err := emailData.GetRandomVerificationCode(ctx, pars.Email)
	if err != nil {
		return err
	}

	// 发送邮件
	err = emailData.SendEmail(ctx, pars.Email, code)
	if err != nil {
		return err
	}

	// 存储验证码
	err = emailData.SaveVerificationCode(ctx, pars.Email, code)
	if err != nil {
		return err
	}

	return nil
}
