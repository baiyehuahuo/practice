DROP TABLE CommentEvents;
CREATE TABLE CommentEvents
(
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '评论 ID',
    `user_id` INT UNSIGNED COMMENT '用户 ID',
    `video_id` INT UNSIGNED COMMENT  '视频 ID',
    `content` VARCHAR(100) COMMENT '评论内容',
    `create_date` VARCHAR(5) COMMENT '评论发布日期，格式 mm-dd',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO CommentEvents (user_id, video_id, content, create_date)
VALUES (1, 1, 'root 测试', '01-01');

INSERT INTO CommentEvents  (`user_id`,  video_id, `content`, `create_date`)
VALUES (2, 2,'抉择之战真不行', '08-11');

INSERT INTO CommentEvents  (`user_id`,  video_id, `content`, `create_date`)
VALUES (2, 2,'抉择之战真垃圾', '08-12');
