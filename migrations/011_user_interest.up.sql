CREATE TABLE IF NOT EXISTS user_interest (
  user_id    BIGINT UNSIGNED NOT NULL COMMENT 'ユーザーID',
  interest_id BIGINT UNSIGNED NOT NULL COMMENT 'インタレスト(興味情報)ID',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY(user_id, interest_id),
  CONSTRAINT user_interest_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_interest_target_id FOREIGN KEY(interest_id) REFERENCES interest(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'userとinterestの紐付きを保存するテーブル'