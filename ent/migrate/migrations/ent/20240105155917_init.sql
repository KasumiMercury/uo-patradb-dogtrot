-- Create "channels" table
CREATE TABLE `channels` (`id` varchar(26) NOT NULL, `display_name` varchar(100) NOT NULL, `channel_id` varchar(26) NOT NULL, `handle` varchar(30) NOT NULL, `thumbnail_url` varchar(140) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `channel_id` (`channel_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "descriptions" table
CREATE TABLE `descriptions` (`id` varchar(26) NOT NULL, `source_id` varchar(12) NOT NULL, `raw` varchar(5000) NOT NULL, `variable` varchar(255) NULL, `template_confidence` bool NOT NULL DEFAULT false, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `periodic_id` varchar(26) NULL, `video_id` varchar(26) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `source_id` (`source_id`), UNIQUE INDEX `video_id` (`video_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "description_changes" table
CREATE TABLE `description_changes` (`id` varchar(26) NOT NULL, `raw` varchar(5000) NOT NULL, `variable` varchar(255) NULL, `changed_at` timestamp NOT NULL, `description_id` varchar(26) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "pat_chats" table
CREATE TABLE `pat_chats` (`id` varchar(26) NOT NULL, `message` varchar(200) NOT NULL, `magnitude` double NOT NULL, `score` double NOT NULL, `is_negative` bool NOT NULL DEFAULT false, `published_at` timestamp NOT NULL, `from_freechat` bool NOT NULL DEFAULT false, `created_at` timestamp NOT NULL, `video_id` varchar(26) NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "periodic_description_templates" table
CREATE TABLE `periodic_description_templates` (`id` varchar(26) NOT NULL, `text` varchar(1000) NOT NULL, `start_use_at` timestamp NULL, `last_use_at` timestamp NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "stream_schedules" table
CREATE TABLE `stream_schedules` (`id` varchar(26) NOT NULL, `scheduled_at` timestamp NOT NULL, `title` varchar(100) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `video_id` varchar(26) NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "videos" table
CREATE TABLE `videos` (`id` varchar(26) NOT NULL, `source_id` varchar(12) NOT NULL, `title` varchar(100) NOT NULL, `duration_seconds` bigint NULL, `is_collaboration` bool NOT NULL DEFAULT false, `status` varchar(255) NOT NULL, `chat_id` varchar(255) NULL, `has_time_range` bool NOT NULL DEFAULT false, `capture_permission` bool NOT NULL DEFAULT true, `scheduled_at` timestamp NULL, `actual_start_at` timestamp NULL, `actual_end_at` timestamp NULL, `published_at` timestamp NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `source_id` (`source_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_disallow_ranges" table
CREATE TABLE `video_disallow_ranges` (`id` varchar(26) NOT NULL, `start_seconds` bigint NOT NULL, `end_seconds` bigint NOT NULL, `video_id` varchar(26) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_play_ranges" table
CREATE TABLE `video_play_ranges` (`id` varchar(26) NOT NULL, `start_seconds` bigint NOT NULL DEFAULT 0, `end_seconds` bigint NULL, `video_id` varchar(26) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_tags" table
CREATE TABLE `video_tags` (`id` varchar(26) NOT NULL, `title` varchar(50) NOT NULL, `normalized_title` varchar(50) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `title` (`title`), UNIQUE INDEX `normalized_title` (`normalized_title`), INDEX `videotag_title` (`title`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_title_changes" table
CREATE TABLE `video_title_changes` (`id` varchar(26) NOT NULL, `title` varchar(100) NOT NULL, `changed_at` timestamp NOT NULL, `video_id` varchar(26) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_channel" table
CREATE TABLE `video_channel` (`video_id` varchar(26) NOT NULL, `channel_id` varchar(26) NOT NULL, PRIMARY KEY (`video_id`, `channel_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "video_tag_videos" table
CREATE TABLE `video_tag_videos` (`video_tag_id` varchar(26) NOT NULL, `video_id` varchar(26) NOT NULL, PRIMARY KEY (`video_tag_id`, `video_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
