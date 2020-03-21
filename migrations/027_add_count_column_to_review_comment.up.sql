ALTER TABLE review_comment
  ADD reply_count    INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'リプライの数' AFTER body,
  ADD favorite_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'お気に入りの数' AFTER reply_count
