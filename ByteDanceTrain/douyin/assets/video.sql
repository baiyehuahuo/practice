DROP TABLE Videos;
CREATE TABLE Videos
(
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '视频唯一标识',
    `author_id` INT UNSIGNED NOT NULL COMMENT '视频作者id',
    `play_url` VARCHAR(64) NOT NULL COMMENT '视频播放地址',
    `cover_url` VARCHAR(64) DEFAULT '' COMMENT '视频封面地址',
    `favorite_count` INT UNSIGNED DEFAULT 0 COMMENT '视频的点赞总数',
    `comment_count` INT UNSIGNED DEFAULT 0 COMMENT '视频的评论总数',
    `is_favorite` BOOLEAN DEFAULT FALSE COMMENT 'true-已点赞，false-未点赞',
    `title` VARCHAR(32) DEFAULT '' COMMENT '视频标题',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO Videos  (`author_id`, `play_url`)
VALUES (1, 'uploadfiles/root/animal.mp4');

INSERT INTO Videos  (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `is_favorite`, `title`)
VALUES (2, 2, 'uploadfiles/fwf/抉择之战.war3', 'uploadfiles/fwf/抉择之战.png', 1, 2, TRUE, '抉择之战');