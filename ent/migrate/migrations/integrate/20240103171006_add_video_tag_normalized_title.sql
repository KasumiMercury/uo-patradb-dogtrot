-- Modify "video_tags" table
ALTER TABLE `video_tags` ADD COLUMN `normalized_title` varchar(50) NOT NULL, ADD UNIQUE INDEX `normalized_title` (`normalized_title`);
