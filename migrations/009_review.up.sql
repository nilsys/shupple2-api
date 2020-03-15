CREATE TABLE IF NOT EXISTS review (
  id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id        BIGINT UNSIGNED NOT NULL COMMENT '投稿したユーザー',
  tourist_spot_id        BIGINT UNSIGNED DEFAULT NULL COMMENT 'tourist_spot_id,inn_idのどちらかが必ず入る',
  inn_id         BIGINT DEFAULT NULL COMMENT '外部APIで使用,tourist_spot_id,inn_idのどちらかが必ず入る',
  score          TINYINT UNSIGNED DEFAULT NULL COMMENT '評価',
  media_count    TINYINT UNSIGNED DEFAULT NULL COMMENT '添付した画像、動画の数',
  body           TEXT NOT NULL COMMENT '本文',
  favorite_count INT UNSIGNED DEFAULT 0 COMMENT 'お気に入りされた数',
  views          BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '閲覧数',
  created_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at     DATETIME DEFAULT NULL,
  PRIMARY KEY(id),
  CONSTRAINT review_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  FULLTEXT KEY(body) WITH PARSER NGRAM,
  CONSTRAINT review_tourist_spot_id FOREIGN KEY(tourist_spot_id) REFERENCES tourist_spot(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS review_media (
  id         CHAR(36) NOT NULL COMMENT 'uuid',
  review_id  BIGINT UNSIGNED NOT NULL COMMENT '紐づくreviewのid',
  priority   INT UNSIGNED NOT NULL COMMENT '何番目の添付ファイルか(1開始)',
  mime_type  VARCHAR(255) NOT NULL COMMENT 'MIME type',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  CONSTRAINT review_media_review_id FOREIGN KEY(review_id) REFERENCES review(id),
  INDEX idx_review_media_review_id_priority (review_id, priority)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS review_hashtag (
  review_id  BIGINT UNSIGNED NOT NULL,
  hashtag_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(review_id, hashtag_id),
  CONSTRAINT review_hashtag_review_id FOREIGN KEY(review_id) REFERENCES review(id),
  CONSTRAINT review_hashtag_hashtag_id FOREIGN KEY(hashtag_id) REFERENCES hashtag(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'reviewに含まれるhashtagを紐付けるためのテーブル';

CREATE TABLE IF NOT EXISTS user_favorite_review (
  user_id BIGINT UNSIGNED NOT NULL,
  review_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, review_id),
  CONSTRAINT review_favorite_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT review_favorite_review_id FOREIGN KEY(review_id) REFERENCES review(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーがお気に入りした投稿';

CREATE TABLE IF NOT EXISTS review_comment (
  id         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id    BIGINT UNSIGNED NOT NULL COMMENT '投稿したユーザー',
  review_id  BIGINT UNSIGNED NOT NULL,
  body       TEXT NOT NULL COMMENT '本文',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  CONSTRAINT review_comment_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT review_comment_review_id FOREIGN KEY(review_id) REFERENCES review(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'レビューに対するコメント';