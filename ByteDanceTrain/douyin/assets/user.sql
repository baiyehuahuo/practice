DROP TABLE Users;
CREATE TABLE Users
(
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '用户 ID',
    `name` VARCHAR(32) NOT NULL UNIQUE COMMENT '用户昵称',
    `password` VARCHAR(32) NOT NULL COMMENT '用户密码',
    `avatar` VARCHAR(64) DEFAULT '' COMMENT '用户头像',
    `background_image` VARCHAR(64) DEFAULT '' COMMENT '用户个人顶部大图',
    `signature` VARCHAR(64) DEFAULT '' COMMENT '用户简介',
    PRIMARY KEY (`id`),
    KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO Users  (`name`, `password`)
VALUES ('root', 'rootpwd');

INSERT INTO Users  (`name`, `password`, `avatar`, `background_image`, `signature`)
VALUES ('user1', 'passwd1', 'uploadfiles/user1/avatar.jpg', 'uploadfiles/user1/background.png', 'hello world');

INSERT INTO Users  (`name`, `password`, `avatar`, `background_image`, `signature`)
VALUES ('user2', 'passwd2', 'uploadfiles/user2/avatar.jpg', 'uploadfiles/user2/background.png', 'fail world');

