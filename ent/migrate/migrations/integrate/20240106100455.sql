-- Create periodic_descriptions index on description table
ALTER TABLE `descriptions` ADD INDEX `periodic_description_idx` (`periodic_id`);

-- Create video_play_ranges index on video_play_ranges table
ALTER TABLE `video_play_ranges` ADD INDEX `video_play_range_idx` (`video_id`);
-- Create video_disallow_ranges index on video_disallow_ranges table
ALTER TABLE `video_disallow_ranges` ADD INDEX `video_disallow_range_idx` (`video_id`);
-- Create video_title_changes index on video_title_changes table
ALTER TABLE `video_title_changes` ADD INDEX `video_title_changes_idx` (`video_id`);
-- Create pat_chats_video index on pat_chats table
ALTER TABLE `pat_chats` ADD INDEX `pat_chats_video_idx` (`video_id`);

-- Create description_changes index on description_changes table
ALTER TABLE `description_changes` ADD INDEX `description_change_description_idx` (`description_id`);

-- Create stream_scheduled_at index stream_schedules table
ALTER TABLE `stream_schedules` ADD INDEX `stream_scheduled_at` (`scheduled_at`);