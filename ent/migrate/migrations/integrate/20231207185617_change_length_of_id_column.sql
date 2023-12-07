-- Modify "category_description_templates" table
ALTER TABLE `category_description_templates` MODIFY COLUMN `id` varchar(26) NOT NULL;
-- Modify "channels" table
ALTER TABLE `channels` MODIFY COLUMN `id` varchar(26) NOT NULL;
-- Modify "description_changes" table
ALTER TABLE `description_changes` MODIFY COLUMN `id` varchar(26) NOT NULL, MODIFY COLUMN `description_id` varchar(26) NOT NULL;
-- Modify "descriptions" table
ALTER TABLE `descriptions` MODIFY COLUMN `id` varchar(26) NOT NULL, MODIFY COLUMN `periodic_id` varchar(26) NULL, MODIFY COLUMN `category_id` varchar(26) NULL, MODIFY COLUMN `video_id` varchar(26) NOT NULL;
-- Modify "pat_chats" table
ALTER TABLE `pat_chats` MODIFY COLUMN `id` varchar(26) NOT NULL, MODIFY COLUMN `video_id` varchar(26) NOT NULL;
-- Modify "periodic_description_templates" table
ALTER TABLE `periodic_description_templates` MODIFY COLUMN `id` varchar(26) NOT NULL;
-- Modify "video_channel" table
ALTER TABLE `video_channel` MODIFY COLUMN `video_id` varchar(26) NOT NULL, MODIFY COLUMN `channel_id` varchar(26) NOT NULL;
-- Modify "video_disallow_ranges" table
ALTER TABLE `video_disallow_ranges` MODIFY COLUMN `id` varchar(26) NOT NULL, MODIFY COLUMN `video_id` varchar(26) NOT NULL;
-- Modify "video_play_ranges" table
ALTER TABLE `video_play_ranges` MODIFY COLUMN `id` varchar(26) NOT NULL, MODIFY COLUMN `video_id` varchar(26) NOT NULL;
-- Modify "video_title_changes" table
ALTER TABLE `video_title_changes` MODIFY COLUMN `id` varchar(26) NOT NULL, MODIFY COLUMN `video_id` varchar(26) NOT NULL;
-- Modify "videos" table
ALTER TABLE `videos` MODIFY COLUMN `id` varchar(26) NOT NULL;
