DROP TABLE RelationEvents;
CREATE TABLE RelationEvents
(
    `user_id` INT UNSIGNED COMMENT '用户 ID',
    `to_user_id` INT UNSIGNED COMMENT '被关注的用户 ID',
    PRIMARY KEY (`user_id`, `to_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO RelationEvents  (`user_id`, `to_user_id`)
VALUES (1, 2);

INSERT INTO RelationEvents  (`user_id`, `to_user_id`)
VALUES (3, 2);

INSERT INTO RelationEvents  (`user_id`, `to_user_id`)
VALUES (2, 3);