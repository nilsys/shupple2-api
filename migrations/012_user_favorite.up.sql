CREATE TABLE IF NOT EXISTS user_favorite_post (
  user_id BIGINT UNSIGNED NOT NULL,
  post_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, post_id),
  CONSTRAINT post_favorite_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT post_favorite_post_id FOREIGN KEY(post_id) REFERENCES post(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーがお気に入りした投稿';

CREATE TABLE IF NOT EXISTS user_favorite_review (
  user_id BIGINT UNSIGNED NOT NULL,
  review_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, review_id),
  CONSTRAINT review_favorite_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT review_favorite_review_id FOREIGN KEY(review_id) REFERENCES review(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーがお気に入りしたレビュー';