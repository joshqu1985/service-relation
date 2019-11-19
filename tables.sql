CREATE TABLE `followers` (
    `id`           BIGINT(20)   NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`          VARCHAR(20)  NOT NULL                COMMENT '',
    `who`          VARCHAR(20)  NOT NULL                COMMENT '粉丝的uid',
    `created_time` BIGINT(20)   NOT NULL DEFAULT 0      COMMENT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_uid_who` (`uid`,`who`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='粉丝表';

CREATE TABLE `followings` (
    `id`           BIGINT(20)   NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`          VARCHAR(20)  NOT NULL                COMMENT '',
    `who`          VARCHAR(20)  NOT NULL                COMMENT '关注的uid',
    `created_time` BIGINT(20)   NOT NULL DEFAULT 0      COMMENT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_uid_who` (`uid`,`who`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='关注表';

CREATE TABLE `mutuals` (
    `id`           BIGINT(20)   NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`          VARCHAR(20)  NOT NULL                COMMENT '',
    `who`          VARCHAR(20)  NOT NULL                COMMENT '关注的uid',
    `created_time` BIGINT(20)   NOT NULL DEFAULT 0      COMMENT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_uid_who` (`uid`,`who`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='相互关注表';

CREATE TABLE `blacks` (
    `id`           BIGINT(20)   NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`          VARCHAR(20)  NOT NULL                COMMENT '',
    `who`          VARCHAR(20)  NOT NULL                COMMENT '关注的uid',
    `created_time` BIGINT(20)   NOT NULL DEFAULT 0      COMMENT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_uid_who` (`uid`,`who`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='黑名单表';

CREATE TABLE `counters` (
    `id`               BIGINT(20)   NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`              VARCHAR(20)  NOT NULL                COMMENT '',
    `followers_count`  INT(11)      NOT NULL DEFAULT 0      COMMENT '粉丝总数',
    `followings_count` INT(11)      NOT NULL DEFAULT 0      COMMENT '关注总数',
    `state`            TINYINT(4)   NOT NULL DEFAULT 0      COMMENT '',
    `created_time`     BIGINT(20)   NOT NULL DEFAULT 0      COMMENT '',
    `updated_time`     BIGINT(20)   NOT NULL DEFAULT 0      COMMENT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_uid` (`uid`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='计数表';
