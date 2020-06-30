ALTER TABLE cf_project ADD favorite_count  INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'いいね数' AFTER support_comment_count;
