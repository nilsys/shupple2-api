CREATE TABLE IF NOT EXISTS hashtag_category (
  category_id  BIGINT UNSIGNED NOT NULL,
  hashtag_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(category_id, hashtag_id),
  CONSTRAINT hashtag_category_category_id FOREIGN KEY(category_id) REFERENCES category(id),
  CONSTRAINT hashtag_category_hashtag_id FOREIGN KEY(hashtag_id) REFERENCES hashtag(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'reviewに含まれるhashtagからhashtagとcategoryを紐付けるためのテーブル';