CREATE TABLE `expenses` (
    `id` CHAR(36) PRIMARY KEY COMMENT '費用ID',
    `name` VARCHAR(255) NOT NULL UNIQUE COMMENT '費目名',
    `created_at` DATETIME(6) NOT NULL COMMENT '作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT '更新日時'
);
