DROP TABLE FavoriteEvents;
CREATE TABLE FavoriteEvents
(
    `user_id` INT UNSIGNED COMMENT '用户 ID',
    `video_id` INT UNSIGNED COMMENT '视频 ID',
    `author_id` INT UNSIGNED COMMENT '视频作者 ID',
    PRIMARY KEY (`user_id`, `video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO FavoriteEvents  (`user_id`, `video_id`, `author_id`)
VALUES (2, 2, 2);

INSERT INTO FavoriteEvents  (`user_id`, `video_id`, `author_id`)
VALUES (2, 3,2);

