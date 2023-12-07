-- Create description_changes index on description_changes table
ALTER TABLE `description_changes` ADD INDEX `description_change_description_idx` (`description_id`);