DROP DATABASE IF EXISTS local_transfer_proj_db;
CREATE DATABASE local_transfer_proj_db; 
USE local_transfer_proj_db;

CREATE TABLE `user_tab` (
    `user_id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `balance` INT UNSIGNED NOT NULL, 
    `created_at` BIGINT UNSIGNED NOT NULL,
    `updated_at` BIGINT UNSIGNED NULL,

    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `transfer_log_tab` (
    `transfer_id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    `user_id_1` BIGINT UNSIGNED NOT NULL,
    `user_id_2` BIGINT UNSIGNED NOT NULL,
    `amount` INT UNSIGNED NOT NULL, 
    `created_at` BIGINT UNSIGNED NOT NULL,
    `status` TINYINT UNSIGNED NOT NULL,

    PRIMARY KEY (`transfer_id`),
    KEY `idx_user_id_1` (`user_id_1`),
    KEY `idx_user_id_2` (`user_id_2`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO local_transfer_proj_db.user_tab
(name, balance, created_at, updated_at)
VALUES('William', 100000, 1669718675585, 1669718675585);

INSERT INTO local_transfer_proj_db.user_tab
(name, balance, created_at, updated_at)
VALUES('Pikachu', 50000, 1669718675585, 1669718675585);

INSERT INTO local_transfer_proj_db.user_tab
(name, balance, created_at, updated_at)
VALUES('Naruto', 50000, 1669718675585, 1669718675585);

INSERT INTO local_transfer_proj_db.user_tab
(name, balance, created_at, updated_at)
VALUES('Ichigo', 50000, 1669718675585, 1669718675585);

INSERT INTO local_transfer_proj_db.user_tab
(name, balance, created_at, updated_at)
VALUES('Luffy', 50000, 1669718675585, 1669718675585);

INSERT INTO local_transfer_proj_db.user_tab
(name, balance, created_at, updated_at)
VALUES('Cid', 50000, 1669718675585, 1669718675585);

INSERT INTO local_transfer_proj_db.user_tab
(name, balance, created_at, updated_at)
VALUES('Aether', 50000, 1669718675585, 1669718675585);

INSERT INTO local_transfer_proj_db.user_tab
(name, balance, created_at, updated_at)
VALUES('Chief', 50000, 1669718675585, 1669718675585);

INSERT INTO local_transfer_proj_db.user_tab
(name, balance, created_at, updated_at)
VALUES('Master', 50000, 1669718675585, 1669718675585);
