-- Modify "pat_chats" table
ALTER TABLE `pat_chats` MODIFY COLUMN `video_id` varchar(26) NULL, ADD COLUMN `from_freechat` bool NOT NULL DEFAULT false;
-- Create "video_tags" table
CREATE TABLE `video_tags` (`id` varchar(26) NOT NULL, `title` varchar(30) NOT NULL, PRIMARY KEY (`id`), INDEX `videotag_title` (`title`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_tag_videos" table
CREATE TABLE `video_tag_videos` (`video_tag_id` varchar(26) NOT NULL, `video_id` varchar(26) NOT NULL, PRIMARY KEY (`video_tag_id`, `video_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
