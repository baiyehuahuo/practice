DROP TABLE MessageEvents;
CREATE TABLE MessageEvents
(
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '消息 ID',
    `to_user_id` INT UNSIGNED COMMENT '消息接受者的 ID',
    `from_user_id` INT UNSIGNED COMMENT '消息发送者的 ID',
    `content` VARCHAR(100) NOT NULL COMMENT '消息内容',
    `create_time` VARCHAR(5) NOT NULL COMMENT '消息创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO MessageEvents  (`to_user_id`, `from_user_id`,`content`, `create_time`)
VALUES (2, 1, '测试', '01-01');

INSERT INTO MessageEvents  (`to_user_id`, `from_user_id`,`content`, `create_time`)
VALUES (2, 3, '开学不快乐', '08-15');
