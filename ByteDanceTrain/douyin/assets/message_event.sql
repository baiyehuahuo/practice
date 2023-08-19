DROP TABLE MessageEvents;
CREATE TABLE MessageEvents
(
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '消息 ID',
    `to_user_id` INT UNSIGNED COMMENT '消息接受者的 ID',
    `from_user_id` INT UNSIGNED COMMENT '消息发送者的 ID',
    `content` VARCHAR(100) NOT NULL COMMENT '消息内容',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '消息创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO MessageEvents  (`to_user_id`, `from_user_id`,`content`, `create_time`)
VALUES (2, 1, '测试', '2023-08-11 01:02:48');

INSERT INTO MessageEvents  (`to_user_id`, `from_user_id`,`content`, `create_time`)
VALUES (2, 3, '开学不快乐', '2023-08-11 01:02:50');

INSERT INTO MessageEvents  (`to_user_id`, `from_user_id`,`content`, `create_time`)
VALUES (3, 2, '早安', '2023-08-11 12:02:48');

INSERT INTO MessageEvents  (`to_user_id`, `from_user_id`,`content`, `create_time`)
VALUES (2, 3, '午安', '2023-08-11 18:02:48');