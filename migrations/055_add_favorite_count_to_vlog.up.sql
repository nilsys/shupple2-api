ALTER TABLE vlog ADD favorite_count  INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'いいね数' AFTER twitter_count;
