CREATE TABLE `book`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `book_name`   varchar(60)  NOT NULL DEFAULT '' COMMENT '书名',
    `author`      varchar(60)  NOT NULL DEFAULT '' COMMENT '作者',
    `category`    varchar(60)  NOT NULL DEFAULT '' COMMENT '分类',
    `location`    varchar(60)  NOT NULL DEFAULT '' COMMENT '现存地点',
    `notice`      varchar(256) NOT NULL DEFAULT '' COMMENT '备注',
    `source`      varchar(256) NOT NULL DEFAULT '' COMMENT '来源',
    `state`       varchar(256) NOT NULL DEFAULT '' COMMENT '状态',
    `keywords`    varchar(256) NOT NULL DEFAULT '' COMMENT '关键字',
    `images`      text COMMENT '图片',
    `abstract`    varchar(256) NOT NULL DEFAULT '' COMMENT '摘要',
    `del`         tinyint      NOT NULL DEFAULT '0' COMMENT '是否已被删除',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
--     UNIQUE KEY `uniq_idx_book_name_author` (`book_name`,`author`)
--     KEY           `idx_name` (`book_name`),
--     KEY           `idx_author` (`author`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='图书'