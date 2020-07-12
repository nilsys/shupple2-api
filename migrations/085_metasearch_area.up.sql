CREATE TABLE IF NOT EXISTS metasearch_area (
  metasearch_area_id   BIGINT UNSIGNED NOT NULL,
  metasearch_area_type TINYINT UNSIGNED NOT NULL,
  area_category_id    BIGINT UNSIGNED NOT NULL,
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (metasearch_area_id, metasearch_area_type),
  CONSTRAINT metasearch_area_category_id FOREIGN KEY(area_category_id) REFERENCES area_category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'メタサーチのエリアとエリアカテゴリの紐付け。N:1';