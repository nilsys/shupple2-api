CREATE TABLE IF NOT EXISTS user_favorite_cf_project (
  user_id       BIGINT UNSIGNED NOT NULL,
  cf_project_id BIGINT UNSIGNED NOT NULL,
  created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, cf_project_id),
  CONSTRAINT user_favorite_cf_project_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_favorite_cf_project_cf_project_id FOREIGN KEY(cf_project_id) REFERENCES cf_project(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーがお気に入りしたクラウドファンディングプロジェクト';
