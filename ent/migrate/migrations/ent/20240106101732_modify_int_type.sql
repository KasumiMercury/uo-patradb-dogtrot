-- Modify "video_tags" table
ALTER TABLE `video_tags` MODIFY COLUMN `series_numbering` tinyint NULL;
-- Modify "videos" table
ALTER TABLE `videos` MODIFY COLUMN `numbering` tinyint NULL;
