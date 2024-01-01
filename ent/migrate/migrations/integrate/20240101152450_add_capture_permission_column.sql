-- Modify "videos" table
ALTER TABLE `videos` ADD COLUMN `capture_permission` bool NOT NULL DEFAULT true;
