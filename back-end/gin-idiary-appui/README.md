<!--
 * @Author: liziwei01
 * @Date: 2022-03-03 15:20:51
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 19:54:21
 * @Description: README
-->
# gin-idiary-appui

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

This is a school project written by senior students Ziwei Li, Kexin Wu from CUHK(SZ)

Interface

post: /email/verificationCode 

get verification code

|getParams|comment|require|
| --------- | --------- | --------- |
|email|user email address|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": null,
    "errno": 0,
    "errmsg": "Success"
}

{
    "data”:”该邮箱60s内已发送过验证码", 
    "errmsg":"Failure",
    "errno”:-1
}
```

post: /upload/image 

upload image

|postParams|comment|require|
| --------- | --------- | --------- |
|user_id|user id|yes|
|file|file|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |
|fileName|file name|yes|

eg
```bash
{
    "data": "123/files/1ae88abb-1c81-4fd9-b0ee-0ade49388058.jpeg",
    "errmsg": "Success",
    "errno": 0
}
```

get: /upload/getImageURL  

get image url

Note that the returned image URL is only valid for a short period

|getParams|comment|require|
| --------- | --------- | --------- |
|file_name|file_name|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |
|file_url|file url|yes|

eg
```bash
{
    "data": "http://idiary-image.oss-cn-shenzhen.aliyuncs.com/123%2Ffiles%2F1ae88abb-1c81-4fd9-b0ee-0ade49388058.jpeg?Expires=1649750074&OSSAccessKeyId=LTAI5tFEUnHRu5htgFXyKjz7&Signature=TTylhQSeXrA083Xk5tIYdos0Vpg%3D",
    "errmsg": "Success",
    "errno": 0
}
```

Mysql require
```bash
DROP table if exists `idiary_diary`;
create table `idiary_diary`(
    `user_id` int NOT NULL,
    `diary_id` int NOT NULL,
    `title` varchar(255) NOT NULL,
    `content` varchar(255) NOT NULL,
    `timestamp` datetime(6) NOT NULL ON UPDATE CURRENT_TIMESTAMP(6),
    `authority` tinyint(1) NOT NULL DEFAULT '0',
    `address` varchar(255) NOT NULL,
    primary key (`diary_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP table if exists `idiary_user`;
create table `idiary_user`(
    `user_id` int NOT NULL,
    `password` varchar(255) NOT NULL,
    `nickname` varchar(255) NOT NULL,
    `city` varchar(255) NOT NULL,
    `picture` varchar(255) DEFAULT NULL COMMENT '',
    primary key (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP table if exists `idiary_friends`;
create table `idiary_friends`(
    `user_id` int NOT NULL,
    `friend_id` int NOT NULL,
    primary key (`user_id`, `friend_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
