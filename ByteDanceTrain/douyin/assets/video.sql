DROP TABLE Videos;
CREATE TABLE Videos
(
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '视频唯一标识',
    `author_id` INT UNSIGNED NOT NULL COMMENT '视频作者id',
    `play_url` VARCHAR(100) NOT NULL COMMENT '视频播放地址',
    `cover_url` VARCHAR(64) DEFAULT '' COMMENT '视频封面地址',
    `title` VARCHAR(32) DEFAULT '' COMMENT '视频标题',
    `publish_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '发表时间',
    PRIMARY KEY (`id`),
    KEY (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO Videos  (`author_id`, `play_url`, `publish_time`)
VALUES (1, 'uploadfiles/root/animal.mp4', '2200-01-04 12:00:00');

# timestamp 1691686968
INSERT INTO Videos  (`id`, `author_id`, `play_url`, `cover_url`, `title`, `publish_time`)
VALUES (2, 2, 'uploadfiles/user1/video_20230603_161305.mp4', 'uploadfiles/user1/WX20230819-143711.png', '长沙西湖公园', '2023-08-11 01:02:48');

# timestamp 1399548023
INSERT INTO Videos  (`id`, `author_id`, `play_url`, `cover_url`, `title`, `publish_time`)
VALUES (3, 2, 'uploadfiles/user1/wx_camera_1675574086089.mp4', 'uploadfiles/user1/WX20230819-143736.png', '莆田瞻圣寺', '2014-05-08 19:20:23');

# timestamp ?
INSERT INTO Videos  (`id`, `author_id`, `play_url`, `cover_url`, `title`, `publish_time`)
VALUES (4, 3, 'uploadfiles/user2/VIDEO_20230819_151710039.mp4', 'uploadfiles/user2/WX20230819-151839.png', '莆田跳火堆', '2013-08-19 15:20:30');