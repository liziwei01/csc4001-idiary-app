/*
 * @Author: liziwei01
 * @Date: 2022-03-21 17:48:04
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-21 18:08:53
 * @Description: file content
 */
package email

import (
	"context"
	"testing"
)

func TestEmail(t *testing.T) {
	ctx := context.Background()
	client, err := GetEmailClient(ctx, "email_idiary_user")
	if err != nil {
		t.Error(err)
	}
	err = client.Send(ctx, "118010160@link.cuhk.edu.cn", "testsub", "testbody")
	if err != nil {
		t.Error(err)
	}
}
