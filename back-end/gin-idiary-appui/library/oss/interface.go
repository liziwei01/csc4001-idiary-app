/*
 * @Author: liziwei01
 * @Date: 2022-03-04 13:52:11
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-04 14:50:10
 * @Description: file content
 */
package oss

import (
	"context"
	"time"
)

type Client interface {
	// GetStr 获取value
	Get(ctx context.Context, key string) (value string, err error)
	// Set 将字符串值 value 关联到 key
	// 如果 key 已经存在， SET 将覆盖旧值 无视类型
	// 过期时间为 nanoseconds 纳秒
	Set(ctx context.Context, key string, value string, expireTime ...time.Duration) error
	// Del
	Del(ctx context.Context, k ...string) error
	// Determine if a key exists
	Exists(ctx context.Context, keys ...string) (bool, error)
}
