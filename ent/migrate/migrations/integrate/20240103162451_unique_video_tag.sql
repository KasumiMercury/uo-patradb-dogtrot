-- Modify "video_tags" table
ALTER TABLE `video_tags` ADD UNIQUE INDEX `title` (`title`);
