CREATE TABLE `book` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `title` varchar(100) DEFAULT NULL,
                        `author` varchar(100) DEFAULT NULL,
                        `publish_date` varchar(100) DEFAULT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;