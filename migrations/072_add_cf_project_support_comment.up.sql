CREATE TABLE IF NOT EXISTS cf_project_support_comment (
  id                 BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id            BIGINT UNSIGNED NOT NULL COMMENT '投稿したユーザーID',
  cf_project_id      BIGINT UNSIGNED NOT NULL COMMENT 'cf_project.id',
  body               TEXT NOT NULL COMMENT '本文',
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  CONSTRAINT cf_project_support_comment FOREIGN KEY(cf_project_id) REFERENCES cf_project(id),
  CONSTRAINT cf_project_support_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'クラウドファンディングプロジェクトに対する応援コメント';