-- Modify "description_changes" table
ALTER TABLE `description_changes` MODIFY COLUMN `raw` varchar(5000) NOT NULL;
-- Modify "descriptions" table
ALTER TABLE `descriptions` MODIFY COLUMN `raw` varchar(5000) NOT NULL;
