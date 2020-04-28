CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstname` varchar(64) NOT NULL DEFAULT '',
  `lastname` varchar(64) NOT NULL DEFAULT '',
  `email` varchar(64) NOT NULL DEFAULT '',
  `password` varchar(256) NOT NULL DEFAULT '',
  `phone` varchar(64) NOT NULL DEFAULT '',
  `signup_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `types` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


CREATE TABLE `options` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `type` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `coffee` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `base_price` DECIMAL(4, 2) NOT NULL DEFAULT 0.00,
  `image` varchar(64),
  `options` varchar(64) DEFAULT NULL,
  `feature` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `tea` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `base_price` DECIMAL(4, 2) NOT NULL DEFAULT 0.00,
  `image` varchar(64),
  `options` varchar(64),
  `feature` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `other_drinks` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `base_price` DECIMAL(4, 2) NOT NULL DEFAULT 0.00,
  `image` varchar(64),
  `options` varchar(64),
  `feature` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `snacks` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `base_price` DECIMAL(4, 2) NOT NULL DEFAULT 0.00,
  `image` varchar(64),
  `options` varchar(64),
  `feature` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `food` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `base_price` DECIMAL(4, 2) NOT NULL DEFAULT 0.00,
  `image` varchar(64),
  `options` varchar(64),
  `feature` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;