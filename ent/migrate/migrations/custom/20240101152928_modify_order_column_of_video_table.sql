-- Modify Order of capture_permission column in video table to be after has_time_range
ALTER TABLE videos MODIFY COLUMN capture_permission TINYINT(1) AFTER has_time_range;