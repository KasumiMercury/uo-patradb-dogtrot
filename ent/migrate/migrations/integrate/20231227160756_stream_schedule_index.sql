-- Create stream_scheduled_at index stream_schedules table
ALTER TABLE `stream_schedules` ADD INDEX `stream_scheduled_at` (`scheduled_at`);