CREATE TABLE IF NOT EXISTS lcategory (
  id         BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  name       VARCHAR(255) NOT NULL COMMENT '名前',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'wordpressのtourist_spot__catに相当する。';

CREATE TABLE IF NOT EXISTS tourist_spot (
  id             BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  slug           VARCHAR(255) NOT NULL COMMENT 'wordpressのslugを入れる',
  name           VARCHAR(255) NOT NULL COMMENT '名前',
  website_url    VARCHAR(1024) DEFAULT NULL COMMENT '公式HP',
  city           VARCHAR(255) NOT NULL COMMENT '県名・都市',
  address        VARCHAR(255) NOT NULL COMMENT '表示用住所',
  lat            DOUBLE NOT NULL COMMENT '緯度',
  lng            DOUBLE NOT NULL COMMENT '経度',
  access_car     VARCHAR(1024) NOT NULL COMMENT '車でのアクセス',
  access_train   VARCHAR(1024) NOT NULL COMMENT '電車でのアクセス',
  access_bus     VARCHAR(1024) DEFAULT NULL COMMENT 'バスでのアクセス',
  opening_hours  VARCHAR(1024) NOT NULL COMMENT '営業時間',
  tel            VARCHAR(255) DEFAULT NULL COMMENT '電話番号',
  price          VARCHAR(1024) NOT NULL COMMENT '価格',
  instagram_url  VARCHAR(255) DEFAULT NULL COMMENT 'instagram',
  search_inn_url VARCHAR(255) DEFAULT NULL COMMENT '関連宿を探すためのstaywayURL',
  created_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at     DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  FULLTEXT KEY(name) WITH PARSER NGRAM,
  UNIQUE INDEX tourist_spot_slug(slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'ほぼwordpressのデータを移しただけなので、もう少し最適化できるはず(search_inn_url等)';

CREATE TABLE IF NOT EXISTS tourist_spot_category (
  tourist_spot_id BIGINT UNSIGNED NOT NULL,
  category_id BIGINT UNSIGNED NOT NULL,
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(tourist_spot_id, category_id),
  CONSTRAINT tourist_spot_category_tourist_spot_id FOREIGN KEY(tourist_spot_id) REFERENCES tourist_spot(id),
  CONSTRAINT tourist_spot_category_category_id FOREIGN KEY(category_id) REFERENCES category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS tourist_spot_lcategory (
  tourist_spot_id  BIGINT UNSIGNED NOT NULL,
  lcategory_id BIGINT UNSIGNED NOT NULL,
  created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(tourist_spot_id, lcategory_id),
  CONSTRAINT tourist_spot_lcategory_tourist_spot_id FOREIGN KEY(tourist_spot_id) REFERENCES tourist_spot(id),
  CONSTRAINT tourist_spot_lcategory_lcategory_id FOREIGN KEY(lcategory_id) REFERENCES lcategory(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
