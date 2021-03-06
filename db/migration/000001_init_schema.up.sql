CREATE TABLE `accounts` (
                            `id` bigint PRIMARY KEY AUTO_INCREMENT,
                            `owner` varchar(255) NOT NULL COMMENT '户主',
                            `balance` bigint NOT NULL,
                            `currency` varchar(255) NOT NULL,
                            `created_at` timestamp NOT NULL DEFAULT current_timestamp
);

CREATE TABLE `entries` (
                           `id` bigint PRIMARY KEY auto_increment,
                           `account_id` bigint NOT NULL,
                           `amount` bigint NOT NULL,
                           `created_at` timestamp NOT NULL DEFAULT current_timestamp
);

CREATE TABLE `transfers` (
                             `id` bigint PRIMARY KEY auto_increment,
                             `from_account_id` bigint NOT NULL,
                             `to_account_id` bigint NOT NULL,
                             `amount` bigint NOT NULL,
                             `created_at` timestamp NOT NULL DEFAULT current_timestamp
);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);

CREATE INDEX `accounts_index_0` ON `accounts` (`owner`);

CREATE INDEX `entries_index_1` ON `entries` (`account_id`);

CREATE INDEX `transfers_index_2` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_4` ON `transfers` (`from_account_id`, `to_account_id`);
