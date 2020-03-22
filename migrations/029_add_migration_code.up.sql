ALTER TABLE user
  ADD migration_code CHAR(36) DEFAULT NULL COMMENT '先にwordpressで登録したユーザーがsign upした時に紐付けるためのキー。UUID' AFTER wordpress_id,
  ADD INDEX idx_user_migration_code(migration_code);
