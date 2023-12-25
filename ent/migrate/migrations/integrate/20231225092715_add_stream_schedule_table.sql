-- Create "stream_schedules" table
CREATE TABLE `stream_schedules` (`id` varchar(26) NOT NULL, `scheduled_at` timestamp NOT NULL, `title` varchar(100) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `video_id` varchar(26) NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
