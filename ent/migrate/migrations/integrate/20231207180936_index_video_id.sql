-- Create video_play_ranges index on video_play_ranges table
ALTER TABLE `video_play_ranges` ADD INDEX `video_play_range_idx` (`video_id`);
-- Create video_disallow_ranges index on video_disallow_ranges table
ALTER TABLE `video_disallow_ranges` ADD INDEX `video_disallow_range_idx` (`video_id`);
-- Create video_title_changes index on video_title_changes table
ALTER TABLE `video_title_changes` ADD INDEX `video_title_changes_idx` (`video_id`);
-- Create pat_chats_video index on pat_chats table
ALTER TABLE `pat_chats` ADD INDEX `pat_chats_video_idx` (`video_id`);