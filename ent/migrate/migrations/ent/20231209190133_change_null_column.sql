-- Modify "videos" table
ALTER TABLE `videos` MODIFY COLUMN `normalized_title` varchar(255) NULL, MODIFY COLUMN `duration_seconds` bigint NULL;
