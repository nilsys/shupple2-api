ALTER TABLE review_comment_reply ADD favorite_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'いいね数' AFTER body;