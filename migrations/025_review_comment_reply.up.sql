CREATE TABLE IF NOT EXISTS review_comment_reply (
  id                 BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id            BIGINT UNSIGNED NOT NULL COMMENT '投稿したユーザーID',
  review_comment_id  BIGINT UNSIGNED NOT NULL COMMENT 'review_comment.id',
  body               TEXT NOT NULL COMMENT '本文',
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  CONSTRAINT review_comment_reply_review_comment_id FOREIGN KEY(review_comment_id) REFERENCES review_comment(id),
  CONSTRAINT review_comment_reply_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'レビューのコメントに対するレビュー';
