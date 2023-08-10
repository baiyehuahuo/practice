DROP TABLE Videos;
CREATE TABLE Videos
(
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '视频唯一标识',
    `author_id` INT UNSIGNED NOT NULL COMMENT '视频作者id',
    `play_url` VARCHAR(64) NOT NULL COMMENT '视频播放地址',
    `cover_url` VARCHAR(64) DEFAULT '' COMMENT '视频封面地址',
    `favorite_count` INT UNSIGNED DEFAULT 0 COMMENT '视频的点赞总数',
    `comment_count` INT UNSIGNED DEFAULT 0 COMMENT '视频的评论总数',
    `is_favorite` BOOL DEFAULT FALSE COMMENT 'true-已点赞，false-未点赞',
    `title` VARCHAR(32) DEFAULT '' COMMENT '视频标题',
    `publish_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '发表时间',
    PRIMARY KEY (`id`),
    KEY (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO Videos  (`author_id`, `play_url`, `publish_time`)
VALUES (1, 'uploadfiles/root/animal.mp4', '2000-01-04 12:00:00');

INSERT INTO Videos  (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `is_favorite`, `title`, `publish_time`)
VALUES (2, 2, 'uploadfiles/fwf/抉择之战.war3', 'uploadfiles/fwf/抉择之战.png', 1, 2, TRUE, '抉择之战 记录视频', '2023-08-11 01:02:48');

INSERT INTO Videos  (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `is_favorite`, `title`, `publish_time`)
VALUES (3, 2, 'uploadfiles/fwf/抉择之战.mp4', 'uploadfiles/fwf/抉择之战.jpg', 3, 4, False, '抉择之战 游戏视频', '2014-05-08 19:20:23');