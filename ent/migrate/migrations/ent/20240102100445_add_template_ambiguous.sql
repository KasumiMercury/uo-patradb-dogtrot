-- Modify "descriptions" table
ALTER TABLE `descriptions` ADD COLUMN `template_confidence` bool NOT NULL DEFAULT true;
-- Modify "videos" table
ALTER TABLE `videos` MODIFY COLUMN `capture_permission` bool NOT NULL DEFAULT true;
