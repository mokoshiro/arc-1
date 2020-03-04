CREATE TABLE `room` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `addr` VARCHAR(32) NOT NULL,
  `location` POINT NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  SPATIAL KEY `location` (`location`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;