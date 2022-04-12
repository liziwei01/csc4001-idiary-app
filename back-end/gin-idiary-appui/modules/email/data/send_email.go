/*
 * @Author: liziwei01
 * @Date: 2022-04-12 13:47:08
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 13:59:20
 * @Description: file content
 */
package data

import (
	"context"
	"gin-idiary-appui/modules/email/constant"
	"gin-idiary-appui/library/email"

)

func SendEmail(ctx context.Context, emailAddress string, code string) error {
	client, err := email.GetClient(ctx, constant.SERVICE_CONF_EMAIL_IDIARY)
	if err != nil {
		return err
	}

	subject := "iDiary Support"
	body := "Hello, your security code is: " + code
	err = client.Send(ctx, emailAddress, subject, body)
	if err != nil {
		return err
	}
	
	return nil
}
