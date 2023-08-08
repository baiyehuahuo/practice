DROP TABLE Users;
CREATE TABLE Users
(
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '用户 ID',
    `name` VARCHAR(32) NOT NULL UNIQUE COMMENT '用户昵称',
    `password` VARCHAR(32) NOT NULL COMMENT '用户密码',
    `follow_count` INT UNSIGNED DEFAULT 0 COMMENT '关注总数',
    `follower_count` INT UNSIGNED DEFAULT 0 COMMENT '粉丝总数',
    `is_follow` BOOLEAN DEFAULT FALSE COMMENT 'true-已关注, false-未关注',
    `avatar` VARCHAR(64) DEFAULT '' COMMENT '用户头像',
    `background_image` VARCHAR(64) DEFAULT '' COMMENT '用户个人顶部大图',
    `signature` VARCHAR(64) DEFAULT '' COMMENT '用户简介',
    `total_favorited` INT DEFAULT 0 COMMENT '获赞数量',
    `work_count` INT DEFAULT  0 COMMENT '作品数量',
    `favorite_count` INT DEFAULT  0 COMMENT '点赞数量',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO Users  (`name`, `password`)
VALUES ('fwf', 'fwf233');

