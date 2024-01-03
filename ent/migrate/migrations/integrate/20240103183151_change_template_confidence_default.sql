-- Modify "descriptions" table
ALTER TABLE `descriptions` MODIFY COLUMN `template_confidence` bool NOT NULL DEFAULT false;
