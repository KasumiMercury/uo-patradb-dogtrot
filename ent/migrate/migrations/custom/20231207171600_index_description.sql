-- Create periodic_descriptions index on description table
ALTER TABLE `descriptions` ADD INDEX `periodic_description_idx` (`periodic_id`);
-- Create category_descriptions index on description table
ALTER TABLE `descriptions` ADD INDEX `category_description_idx` (`category_id`);