CREATE TABLE `users` (
                         `username` varchar(255) PRIMARY KEY,
                         `hashed_password` varchar(255) NOT NULL,
                         `full_name` varchar(255) NOT NULL,
                         `email` varchar(255) UNIQUE NOT NULL,
                         `password_changed_at` timestamp DEFAULT current_timestamp,
                         `created_at` timestamp DEFAULT current_timestamp
);

ALTER TABLE `accounts` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`username`);

CREATE UNIQUE INDEX `owner_currency_key` ON `accounts` (`owner`, `currency`);

# ALTER TABLE `accounts` ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner","currency");
