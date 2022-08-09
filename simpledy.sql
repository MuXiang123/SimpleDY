SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`             bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '用户id，自增主键',
    `name`           varchar(255) NOT NULL COMMENT '昵称',
    `username`       varchar(255) NOT NULL COMMENT '用户名',
    `password`       varchar(255) NOT NULL COMMENT '密码',
    `follow_count`   bigint(20) COMMENT '关注数',
    `follower_count` bigint(20) COMMENT '粉丝数',
    PRIMARY KEY (`id`),
    KEY `name_password_idx` (`name`, `password`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8 COMMENT ='用户表';

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like`
(
    `id`       bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`  bigint(20) NOT NULL COMMENT '点赞用户id',
    `video_id` bigint(20) NOT NULL COMMENT '被点赞的视频id',
    `cancel`   tinyint(4) NOT NULL DEFAULT '0' COMMENT '默认点赞为0，取消赞为1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `userIdtoVideoIdIdx` (`user_id`,`video_id`) USING BTREE,
    KEY        `userIdIdx` (`user_id`) USING BTREE,
    KEY        `videoIdx` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1229 DEFAULT CHARSET=utf8 COMMENT='点赞表';

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`     bigint(20) NOT NULL COMMENT '用户id',
    `follower_id` bigint(20) NOT NULL COMMENT '关注的用户',
    PRIMARY KEY (`id`),
    UNIQUE KEY `userIdToFollowerIdIdx` (`user_id`,`follower_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='关注表';

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键，视频唯一id',
    `author_id`    bigint(20) NOT NULL COMMENT '视频作者id',
    `video_path`     varchar(255) NOT NULL COMMENT '播放url',
    `cover_path`    varchar(255) NOT NULL COMMENT '封面url',
    `favorite_count` bigint(20) COMMENT '点赞数',
    `comment_count` bigint(20) COMMENT '评论数',
    `create_at` datetime     NOT NULL COMMENT '发布时间',
    `title`        varchar(255) DEFAULT NULL COMMENT '视频名称',
    PRIMARY KEY (`id`),
    KEY            `time` (`create_at`) USING BTREE,
    KEY            `author` (`author_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='视频表';

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id，自增主键',
    `user_id`      bigint(20) NOT NULL COMMENT '评论发布用户id',
    `video_id`     bigint(20) NOT NULL COMMENT '评论视频id',
    `comment_text` varchar(255) NOT NULL COMMENT '评论内容',
    `create_date`  datetime     NOT NULL COMMENT '评论发布时间',
    `cancel`       tinyint(4) NOT NULL DEFAULT '0' COMMENT '默认评论发布为0，取消后为1',
    PRIMARY KEY (`id`),
    KEY            `videoIdIdx` (`video_id`) USING BTREE COMMENT '评论列表使用视频id作为索引-方便查看视频下的评论列表'
) ENGINE=InnoDB AUTO_INCREMENT=1206 DEFAULT CHARSET=utf8 COMMENT='评论表';