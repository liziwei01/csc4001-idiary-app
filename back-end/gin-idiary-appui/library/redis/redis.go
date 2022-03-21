/*
 * @Author: liziwei01
 * @Date: 2022-03-21 22:36:04
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-21 22:39:31
 * @Description: file content
 */
package redis

import (
	"context"
	"time"
)

func (c *client) Get(ctx context.Context, key string) (value string, err error) {
	return "", nil
}

func (c *client) Set(ctx context.Context, key string, value string, expireTime ...time.Duration) error {
	return nil
}

func (c *client) Del(ctx context.Context, k ...string) error {
	return nil
}

func (c *client) Exists(ctx context.Context, keys ...string) (bool, error) {
	return false, nil
}
