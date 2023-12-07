-- Create "category_description_templates" table
CREATE TABLE `category_description_templates` (`id` varchar(255) NOT NULL, `text` varchar(255) NOT NULL, `start_use_at` timestamp NULL, `last_use_at` timestamp NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "channels" table
CREATE TABLE `channels` (`id` varchar(255) NOT NULL, `display_name` varchar(255) NOT NULL, `channel_id` varchar(255) NOT NULL, `handle` varchar(255) NOT NULL, `thumbnail_url` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `channel_id` (`channel_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "descriptions" table
CREATE TABLE `descriptions` (`id` varchar(255) NOT NULL, `raw` varchar(255) NOT NULL, `variable` varchar(255) NULL, `normalized_variable` varchar(255) NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `periodic_id` varchar(255) NULL, `category_id` varchar(255) NULL, `video_id` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `video_id` (`video_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "description_changes" table
CREATE TABLE `description_changes` (`id` varchar(255) NOT NULL, `raw` varchar(255) NOT NULL, `variable` varchar(255) NULL, `normalized_variable` varchar(255) NULL, `changed_at` timestamp NOT NULL, `description_id` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "pat_chats" table
CREATE TABLE `pat_chats` (`id` varchar(255) NOT NULL, `message` varchar(255) NOT NULL, `magnitude` double NOT NULL, `score` double NOT NULL, `is_negative` bool NOT NULL DEFAULT false, `published_at` timestamp NOT NULL, `created_at` timestamp NOT NULL, `video_id` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "periodic_description_templates" table
CREATE TABLE `periodic_description_templates` (`id` varchar(255) NOT NULL, `text` varchar(255) NOT NULL, `start_use_at` timestamp NULL, `last_use_at` timestamp NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "videos" table
CREATE TABLE `videos` (`id` varchar(255) NOT NULL, `video_id` varchar(255) NOT NULL, `title` varchar(255) NOT NULL, `normalized_title` varchar(255) NOT NULL, `duration_seconds` bigint NOT NULL, `is_collaboration` bool NOT NULL, `status` varchar(255) NOT NULL, `chat_id` varchar(255) NULL, `has_time_range` bool NOT NULL DEFAULT false, `scheduled_at` timestamp NULL, `actual_start_at` timestamp NULL, `published_at` timestamp NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `video_id` (`video_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_disallow_ranges" table
CREATE TABLE `video_disallow_ranges` (`id` varchar(255) NOT NULL, `start_seconds` bigint NOT NULL, `end_seconds` bigint NOT NULL, `video_id` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_play_ranges" table
CREATE TABLE `video_play_ranges` (`id` varchar(255) NOT NULL, `start_seconds` bigint NOT NULL DEFAULT 0, `end_seconds` bigint NULL, `video_id` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_title_changes" table
CREATE TABLE `video_title_changes` (`id` varchar(255) NOT NULL, `title` varchar(255) NOT NULL, `normalized_title` varchar(255) NOT NULL, `changed_at` timestamp NOT NULL, `video_id` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_channel" table
CREATE TABLE `video_channel` (`video_id` varchar(255) NOT NULL, `channel_id` varchar(255) NOT NULL, PRIMARY KEY (`video_id`, `channel_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
