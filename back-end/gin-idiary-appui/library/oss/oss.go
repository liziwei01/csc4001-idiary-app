/*
 * @Author: liziwei01
 * @Date: 2022-03-20 18:17:39
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-20 19:41:05
 * @Description: file content
 */
package oss

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func (c *client) Get(ctx context.Context, bucket string, objectKey string) (*bytes.Reader, error) {
	ossBucket, err := c.connect(ctx, bucket)
	if err != nil {
		return nil, err
	}
	file, err := ossBucket.GetObject(objectKey)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(fileBytes), nil
}

func (c *client) Put(ctx context.Context, bucket string, objectKey string, fileReader *bytes.Reader) error {
	ossBucket, err := c.connect(ctx, bucket)
	if err != nil {
		return err
	}
	err = ossBucket.PutObject(objectKey, fileReader)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Del(ctx context.Context, bucket string, objectKey string) error {
	ossBucket, err := c.connect(ctx, bucket)
	if err != nil {
		return err
	}
	return ossBucket.DeleteObject(objectKey)
}

func (c *client) GetURL(ctx context.Context, bucket string, objectKey string) (string, error) {
	ossBucket, err := c.connect(ctx, bucket)
	if err != nil {
		return "", err
	}
	return ossBucket.SignURL(objectKey, oss.HTTPGet, 60)
}