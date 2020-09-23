CREATE TABLE IF NOT EXISTS user_block_user (
  user_id         BIGINT UNSIGNED NOT NULL COMMENT 'ブロックしたユーザー.id',
  blocked_user_id BIGINT UNSIGNED NOT NULL COMMENT 'ブロックされるユーザー.id',
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, blocked_user_id),
  CONSTRAINT user_block_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_block_blocked_user_id FOREIGN KEY(blocked_user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーがユーザーをブロックしたコマンドを表現';