create table IF not exists `climbed_mountains`
(
    `id`                INT(20) AUTO_INCREMENT,
    `name`              VARCHAR(20) NOT NULL,
    `height`            INT(20),
    `climbed_date`      DATE NOT NULL,
    `weather`           VARCHAR(20) NOT NULL,
    `number`            INT(20),
    `description`       VARCHAR(255) NOT NULL,
    `image_url`         VARCHAR(255),
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
