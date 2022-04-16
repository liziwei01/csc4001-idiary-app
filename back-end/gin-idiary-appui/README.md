<!--
 * @Author: liziwei01
 * @Date: 2022-03-03 15:20:51
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 18:34:51
 * @Description: README
-->
# gin-idiary-appui

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

This is a school project written by senior students Ziwei Li, Kexin Wu from CUHK(SZ)

Interface

post: /user/follow/follow

a user want to follow somebody

|postParams|comment|require|
| --------- | --------- | --------- |
|user_id|user ID|yes|
|following_id|the man who's gonna be followed|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": null,
    "errno": 0,
    "errmsg": "Success"
}
```

get: /user/follow/followings

a user's following list

|getParams|comment|require|
| --------- | --------- | --------- |
|user_id|user ID|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": [1, 2, 3],
    "errno": 0,
    "errmsg": "Success"
}
```

get: /user/follow/followers

a user's follower list

|getParams|comment|require|
| --------- | --------- | --------- |
|user_id|user ID|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": [1, 2, 3],
    "errno": 0,
    "errmsg": "Success"
}
```

post: /email/verificationCode 

get verification code

|postParams|comment|require|
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
In database db_idiary_feed:

DROP table if exists `tb_user_diary_feed`;
create table `tb_user_diary_feed`(
    `diary_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'auto increment primary key',
    `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送用户ID',
    `content` varchar(2000) NOT NULL DEFAULT '' COMMENT '投稿内容',
    `image_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '投稿内包含的图片,使用 json list ["image_name1", "image_name2"...]保存',
    `device` varchar(24) NOT NULL DEFAULT '' COMMENT '设备',
    `db_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `authority` tinyint(1) NOT NULL DEFAULT '0',
    `address` varchar(24) NOT NULL DEFAULT '' COMMENT '地址',
    `vote_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `dislike_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `share_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分享数量',
    `report_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '举报数量',
    `delete_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态,0=正常',
    `tags` varchar(255) NOT NULL DEFAULT '' COMMENT 'AI 标签',
    PRIMARY KEY (`diary_id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户日记投稿表';

In database db_idiary_user:

DROP table if exists `tb_user_private_info`;
create table `tb_user_private_info`(
    `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `password` varchar(24) NOT NULL DEFAULT '' COMMENT '用户密码',
    `nickname` varchar(24) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `address` varchar(24) NOT NULL DEFAULT '' COMMENT '地址',
    `profile` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
    `db_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    primary key (`user_id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户隐私信息表';

DROP table if exists `tb_user_followings`;
create table `tb_user_followings`(
    `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `following_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '投稿内包含的图片,使用 json list [user_id1, user_id2...]保存',
    primary key (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户关注列表';

DROP table if exists `tb_user_followers`;
create table `tb_user_followers`(
    `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `follower_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '投稿内包含的图片,使用 json list [user_id1, user_id2...]保存',
    primary key (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户粉丝列表';
```
