CREATE TABLE `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(25) NOT NULL,
    `email` varchar(50) NOT NULL,
    `image` text NOT NULL,
    `account` tinyint(1) NOT NULL,
    `count_login` tinyint(1) NOT NULL DEFAULT 1,
    `registered_at` datetime NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci