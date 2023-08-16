DROP TABLE Users;
CREATE TABLE Users
(
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '用户 ID',
    `name` VARCHAR(32) NOT NULL UNIQUE COMMENT '用户昵称',
    `password` VARCHAR(32) NOT NULL COMMENT '用户密码',
    `is_follow` BOOLEAN DEFAULT FALSE COMMENT 'true-已关注, false-未关注',
    `avatar` VARCHAR(64) DEFAULT '' COMMENT '用户头像',
    `background_image` VARCHAR(64) DEFAULT '' COMMENT '用户个人顶部大图',
    `signature` VARCHAR(64) DEFAULT '' COMMENT '用户简介',
    PRIMARY KEY (`id`),
    KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO Users  (`name`, `password`)
VALUES ('root', 'rootpwd');

INSERT INTO Users  (`name`, `password`, `is_follow`, `avatar`, `background_image`, `signature`)
VALUES ('fwf', 'fwf233', false, '/uploadfiles/fwf/avatar.png', '/uploadfiles/fwf/background.png', 'hello world');

INSERT INTO Users  (`name`, `password`, `is_follow`, `avatar`, `background_image`, `signature`)
VALUES ('范伟锋', 'baiyehuahuo', false, '/uploadfiles/范伟锋/avatar.png', '/uploadfiles/范伟锋/background.png', 'fail world');

