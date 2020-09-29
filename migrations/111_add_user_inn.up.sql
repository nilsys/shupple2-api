CREATE TABLE IF NOT EXISTS user_inn (
  user_id         BIGINT UNSIGNED NOT NULL COMMENT 'user.id',
  inn_id          BIGINT UNSIGNED NOT NULL COMMENT 'inn.id',
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, inn_id),
  CONSTRAINT user_inn_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーと関係する（オーナー等）inn';