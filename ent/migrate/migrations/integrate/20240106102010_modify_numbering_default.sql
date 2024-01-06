-- Modify "video_tags" table
ALTER TABLE `video_tags` MODIFY COLUMN `series_numbering` tinyint NOT NULL DEFAULT 1;
-- Modify "videos" table
ALTER TABLE `videos` MODIFY COLUMN `numbering` tinyint NOT NULL DEFAULT 1;
