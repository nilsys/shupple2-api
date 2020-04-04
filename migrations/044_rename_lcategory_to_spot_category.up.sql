DROP TABLE IF EXISTS tourist_spot_lcategory;
DROP TABLE IF EXISTS lcategory;

CREATE TABLE IF NOT EXISTS spot_category (
  id                         BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  name                       VARCHAR(255) NOT NULL COMMENT '名前',
  slug                       VARCHAR(255) NOT NULL COMMENT 'スラグ',
  type                       TINYINT UNSIGNED NOT NULL COMMENT 'スポットカテゴリ種別。スポットorサブスポット',
  spot_category_id           BIGINT UNSIGNED NOT NULL COMMENT '検索用に非正規化されたスポットカテゴリID',
  sub_spot_category_id       BIGINT UNSIGNED DEFAULT NULL COMMENT '検索用に非正規化されたサブスポットカテゴリID',
  created_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at                 DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT spot_category_spot_category_id FOREIGN KEY(spot_category_id) REFERENCES spot_category(id),
  CONSTRAINT spot_category_sub_spot_category_id FOREIGN KEY(sub_spot_category_id) REFERENCES spot_category(id),
  INDEX idx_spot_category_slug(slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'wordpressのlocation__catに相当する。';

CREATE TABLE IF NOT EXISTS tourist_spot_spot_category (
  tourist_spot_id  BIGINT UNSIGNED NOT NULL,
  spot_category_id BIGINT UNSIGNED NOT NULL,
  created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(tourist_spot_id, spot_category_id),
  CONSTRAINT tourist_spot_spot_category_tourist_spot_id FOREIGN KEY(tourist_spot_id) REFERENCES tourist_spot(id),
  CONSTRAINT tourist_spot_spot_category_spot_category_id FOREIGN KEY(spot_category_id) REFERENCES spot_category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
