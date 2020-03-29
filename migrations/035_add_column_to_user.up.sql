ALTER TABLE user
  ADD header_uuid CHAR(36) DEFAULT NULL COMMENT 's3上のヘッダー画像のUUID' AFTER avatar_uuid,
  ALTER avatar_uuid DROP DEFAULT
