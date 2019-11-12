CREATE TABLE `peer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `peer_id` VARCHAR(512) NOT NULL UNIQUE,
  `addr` VARCHAR(32) NOT NULL,
  `location` POINT NOT NULL,
  `h3_hash` VARCHAR(32) NOT NULL,
  `h3_resolution` int(11),
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  SPATIAL KEY `location` (`location`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;