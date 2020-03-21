CREATE TABLE IF NOT EXISTS user_favorite_review_comment (
  user_id           BIGINT UNSIGNED NOT NULL,
  review_comment_id BIGINT UNSIGNED NOT NULL,
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, review_comment_id),
  CONSTRAINT user_favorite_review_comment_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_favorite_review_comment_review_comment_id FOREIGN KEY(review_comment_id) REFERENCES review_comment(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーがお気に入りしたレビューのコメント';