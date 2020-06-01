CREATE TABLE IF NOT EXISTS user_favorite_comic (
  user_id  BIGINT UNSIGNED NOT NULL,
  comic_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, comic_id),
  CONSTRAINT user_favorite_comic_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_favorite_comic_comic_id FOREIGN KEY(comic_id) REFERENCES comic(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ユーザーがお気に入りした漫画';
