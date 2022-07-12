CREATE TABLE `currencies` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(255) NOT NULL,
  `value` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `currency_request_id` int NOT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `currency_requests` (
  `id` int NOT NULL AUTO_INCREMENT,
  `request_at` timestamp NOT NULL,
  `duration_time_seconds` int NOT NULL,
  PRIMARY KEY (`id`)
)