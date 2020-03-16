CREATE TABLE IF NOT EXISTS feature (
  id          BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  user_id     BIGINT UNSIGNED NOT NULL COMMENT '執筆者',
  slug        VARCHAR(255) NOT NULL COMMENT 'wordpressのslugを入れる',
  title       VARCHAR(1024) NOT NULL COMMENT '記事タイトル',
  body        LONGTEXT NOT NULL COMMENT '本文',
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at  DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX feature_slug(slug),
  CONSTRAINT feature_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS feature_post (
  feature_id BIGINT UNSIGNED NOT NULL,
  post_id    BIGINT UNSIGNED NOT NULL,
  PRIMARY KEY(feature_id, post_id),
  CONSTRAINT feature_post_feature_id FOREIGN KEY(feature_id) REFERENCES feature(id),
  CONSTRAINT feature_post_post_id FOREIGN KEY(post_id) REFERENCES post(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'featureに紐づくpostを保存するテーブル';
