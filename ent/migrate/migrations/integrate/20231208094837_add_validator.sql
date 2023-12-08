-- Modify "channels" table
ALTER TABLE `channels` MODIFY COLUMN `display_name` varchar(100) NOT NULL, MODIFY COLUMN `channel_id` varchar(26) NOT NULL, MODIFY COLUMN `handle` varchar(30) NOT NULL, MODIFY COLUMN `thumbnail_url` varchar(140) NOT NULL;
-- Modify "description_changes" table
ALTER TABLE `description_changes` MODIFY COLUMN `raw` varchar(1000) NOT NULL;
-- Modify "descriptions" table
ALTER TABLE `descriptions` MODIFY COLUMN `raw` varchar(1000) NOT NULL;
-- Modify "pat_chats" table
ALTER TABLE `pat_chats` MODIFY COLUMN `message` varchar(200) NOT NULL;
-- Modify "periodic_description_templates" table
ALTER TABLE `periodic_description_templates` MODIFY COLUMN `text` varchar(1000) NOT NULL;
-- Modify "video_title_changes" table
ALTER TABLE `video_title_changes` MODIFY COLUMN `title` varchar(100) NOT NULL;
-- Modify "videos" table
ALTER TABLE `videos` MODIFY COLUMN `video_id` varchar(12) NOT NULL, MODIFY COLUMN `title` varchar(100) NOT NULL, MODIFY COLUMN `is_collaboration` bool NOT NULL DEFAULT false;
