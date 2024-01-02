-- Modify Order of from_freechat column in pat_chat table to be after published_at
ALTER TABLE pat_chats MODIFY COLUMN from_freechat TINYINT(1) NOT NULL DEFAULT 0 AFTER published_at;