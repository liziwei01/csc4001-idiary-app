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
    `picture` varchar(255) DEFAULT NULL COMMENT '头像，不知道用什么格式'
    primary key (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP table if exists `idiary_friends`;
create table `idiary_friends`(
    `user_id` int NOT NULL,
    `friend_id` int NOT NULL,
    primary key (`user_id`, `friend_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;