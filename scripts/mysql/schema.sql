CREATE TABLE `user_tab` (
    `user_id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `balance` INT UNSIGNED NOT NULL, 
    `created_at` BIGINT UNSIGNED NOT NULL,
    `updated_at` BIGINT UNSIGNED NULL

    PRIMARY KEY (`user_id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `transfer_log_tab` (
    `transfer_id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    `user_id_1` BIGINT UNSIGNED NOT NULL,
    `user_id_1` BIGINT UNSIGNED NOT NULL,
    `amount` INT UNSIGNED NOT NULL, 
    `created_at` BIGINT UNSIGNED NOT NULL,
    `status` TINYINT UNSIGNED NOT NULL

    PRIMARY KEY (`transfer_id`),
    KEY `idx_user_id_1` (`user_id_1`)
    KEY `idx_user_id_2` (`user_id_2`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
