CREATE DATABASE goLottery;

USE goLottery;

CREATE TABLE IF NOT EXISTS `gift` (
    `id` INT AUTO_INCREMENT,
    `title` VARCHAR(255),
    `prize_num` INT,
    `left_num` INT,
    `prize_code` VARCHAR(50),
    `prize_time` INT,
    `img` VARCHAR(255),
    `display_order` INT,
    `gtype` INT,
    `gdata` VARCHAR(255),
    `time_begin` DATETIME,
    `time_end` DATETIME,
    `prize_data` mediumtext,
    `prize_begin` DATETIME,
    `prize_end` DATETIME,
    `sys_status` SMALLINT,
    `sys_created` DATETIME,
    `sys_updated` DATETIME,
    `sys_ip` VARCHAR(50),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `code` (
    `id` INT AUTO_INCREMENT,
    `gift_id` INT,
    `code` VARCHAR(255),
    `sys_created` DATETIME,
    `sys_updated` DATETIME,
    `sys_status` SMALLINT,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`gift_id`) REFERENCES `gift` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `result` (
    `id` INT AUTO_INCREMENT,
    `gift_id` INT,
    `gift_name` VARCHAR(250),
    `gift_type` INT,
    `uid` INT,
    `username` VARCHAR(50),
    `prize_code` INT,
    `gift_data` VARCHAR(50),
    `sys_created` DATETIME,
    `sys_ip` VARCHAR(50),
    `sys_status` SMALLINT,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`gift_id`) REFERENCES `gift` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `black_user` (
    `id` INT AUTO_INCREMENT,
    `username` VARCHAR(50),
    `black_time` DATETIME,
    `real_name` VARCHAR(50),
    `mobile` VARCHAR(50),
    `address` VARCHAR(255),
    `sys_created` DATETIME,
    `sys_updated` DATETIME,
    `sys_ip` VARCHAR(50),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `black_ip` (
    `id` INT AUTO_INCREMENT,
    `ip` VARCHAR(50),
    `black_time` DATETIME,
    `sys_created` DATETIME,
    `sys_updated` DATETIME,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `user_day` (
    `id` INT AUTO_INCREMENT,
    `uid` INT,
    `day` VARCHAR(8),
    `num` INT,
    `sys_created` DATETIME,
    `sys_updated` DATETIME,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


