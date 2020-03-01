CREATE TABLE IF NOT EXISTS user_followee (
  user_id    BIGINT UNSIGNED NOT NULL COMMENT 'フォローした人',
  target_id BIGINT UNSIGNED NOT NULL COMMENT 'フォローされる人',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY(user_id, target_id),
  CONSTRAINT user_follow_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_follow_target_id FOREIGN KEY(target_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'フォロー「した」コマンドを保存するテーブル';

CREATE TABLE IF NOT EXISTS user_followed (
  user_id    BIGINT UNSIGNED NOT NULL COMMENT 'フォローされる人',
  target_id BIGINT UNSIGNED NOT NULL COMMENT 'フォローした人',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY(user_id, target_id),
  CONSTRAINT user_followed_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_followed_target_id FOREIGN KEY(target_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'フォロー「された」コマンドを保存するテーブル';

CREATE TABLE IF NOT EXISTS user_follow_hashtag (
  user_id    BIGINT UNSIGNED NOT NULL COMMENT 'ユーザーID',
  hashtag_id BIGINT UNSIGNED NOT NULL COMMENT 'ハッシュタグID',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY(user_id, hashtag_id),
  CONSTRAINT user_follow_hashtag_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_follow_hashtag_hashtag_id FOREIGN KEY(hashtag_id) REFERENCES hashtag(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーがハッシュタグをフォローしたコマンドを保存するテーブル';