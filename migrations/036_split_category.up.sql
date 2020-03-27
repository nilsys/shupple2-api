CREATE TABLE IF NOT EXISTS area_category (
  id                         BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  name                       VARCHAR(255) NOT NULL COMMENT '名前',
  slug                       VARCHAR(255) NOT NULL COMMENT 'スラグ',
  type                       TINYINT UNSIGNED NOT NULL COMMENT 'エリアカテゴリ種別。エリアorサブエリアorサブサブエリア',
  area_group                 TINYINT UNSIGNED NOT NULL COMMENT 'エリアグループ',
  area_id                    BIGINT UNSIGNED NOT NULL COMMENT '検索用に非正規化されたエリアID',
  sub_area_id                BIGINT UNSIGNED DEFAULT NULL COMMENT '検索用に非正規化されたサブエリアID',
  sub_sub_area_id            BIGINT UNSIGNED DEFAULT NULL COMMENT '検索用に非正規化されたサブサブエリアID',
  metasearch_area_id         BIGINT UNSIGNED COMMENT '宿泊検索側のarea_id',
  metasearch_sub_area_id     BIGINT UNSIGNED COMMENT '宿泊検索側のsub_area_id',
  metasearch_sub_sub_area_id BIGINT UNSIGNED COMMENT '宿泊検索側のsub_sub_area_id',
  created_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at                 DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT area_category_area_id FOREIGN KEY(area_id) REFERENCES area_category(id),
  CONSTRAINT area_category_sub_area_id FOREIGN KEY(sub_area_id) REFERENCES area_category(id),
  CONSTRAINT area_category_sub_sub_area_id FOREIGN KEY(sub_area_id) REFERENCES area_category(id),
  INDEX idx_area_category_area_group(area_group),
  INDEX idx_area_category_slug(slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'wordpressのcategoryのうちエリアに関するもの';

CREATE TABLE IF NOT EXISTS theme_category (
  id                         BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  name                       VARCHAR(255) NOT NULL COMMENT '名前',
  slug                       VARCHAR(255) NOT NULL COMMENT 'スラグ',
  type                       TINYINT UNSIGNED NOT NULL COMMENT 'テーマカテゴリ種別。テーマorサブテーマ',
  theme_id                   BIGINT UNSIGNED NOT NULL COMMENT '検索用に非正規化されたテーマID',
  sub_theme_id               BIGINT UNSIGNED DEFAULT NULL COMMENT '検索用に非正規化されたサブテーマID',
  metasearch_inn_type_id     BIGINT UNSIGNED COMMENT '宿泊検索側の宿タイプid',
  created_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at                 DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT theme_category_theme_id FOREIGN KEY(theme_id) REFERENCES theme_category(id),
  CONSTRAINT theme_category_sub_theme_id FOREIGN KEY(sub_theme_id) REFERENCES theme_category(id),
  INDEX idx_theme_category_slug(slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'wordpressのcategoryのうちテーマに関するもの';

CREATE TABLE IF NOT EXISTS post_area_category (
  post_id          BIGINT UNSIGNED NOT NULL,
  area_category_id BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのarea_categoryに相当する',
  created_at       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(post_id, area_category_id),
  CONSTRAINT post_area_category_post_id FOREIGN KEY(post_id) REFERENCES post(id),
  CONSTRAINT post_area_category_area_category_id FOREIGN KEY(area_category_id) REFERENCES area_category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'postに紐づくarea_categoryを保存するテーブル';

CREATE TABLE IF NOT EXISTS post_theme_category (
  post_id           BIGINT UNSIGNED NOT NULL,
  theme_category_id BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのtheme_categoryに相当する',
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(post_id, theme_category_id),
  CONSTRAINT post_theme_category_post_id FOREIGN KEY(post_id) REFERENCES post(id),
  CONSTRAINT post_theme_category_theme_category_id FOREIGN KEY(theme_category_id) REFERENCES theme_category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'postに紐づくtheme_categoryを保存するテーブル';

CREATE TABLE IF NOT EXISTS tourist_spot_area_category (
  tourist_spot_id  BIGINT UNSIGNED NOT NULL,
  area_category_id BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのarea_categoryに相当する',
  created_at       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(tourist_spot_id, area_category_id),
  CONSTRAINT tourist_spot_area_category_tourist_spot_id FOREIGN KEY(tourist_spot_id) REFERENCES tourist_spot(id),
  CONSTRAINT tourist_spot_area_category_area_category_id FOREIGN KEY(area_category_id) REFERENCES area_category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'tourist_spotに紐づくarea_categoryを保存するテーブル';

CREATE TABLE IF NOT EXISTS tourist_spot_theme_category (
  tourist_spot_id   BIGINT UNSIGNED NOT NULL,
  theme_category_id BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのtheme_categoryに相当する',
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(tourist_spot_id, theme_category_id),
  CONSTRAINT tourist_spot_theme_category_tourist_spot_id FOREIGN KEY(tourist_spot_id) REFERENCES tourist_spot(id),
  CONSTRAINT tourist_spot_theme_category_theme_category_id FOREIGN KEY(theme_category_id) REFERENCES theme_category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'tourist_spotに紐づくtheme_categoryを保存するテーブル';

CREATE TABLE IF NOT EXISTS vlog_area_category (
  vlog_id          BIGINT UNSIGNED NOT NULL,
  area_category_id BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのarea_categoryに相当する',
  created_at       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(vlog_id, area_category_id),
  CONSTRAINT vlog_area_category_vlog_id FOREIGN KEY(vlog_id) REFERENCES vlog(id),
  CONSTRAINT vlog_area_category_area_category_id FOREIGN KEY(area_category_id) REFERENCES area_category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'vlogに紐づくarea_categoryを保存するテーブル';

CREATE TABLE IF NOT EXISTS vlog_theme_category (
  vlog_id           BIGINT UNSIGNED NOT NULL,
  theme_category_id BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのtheme_categoryに相当する',
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(vlog_id, theme_category_id),
  CONSTRAINT vlog_theme_category_vlog_id FOREIGN KEY(vlog_id) REFERENCES vlog(id),
  CONSTRAINT vlog_theme_category_theme_category_id FOREIGN KEY(theme_category_id) REFERENCES theme_category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'vlogに紐づくtheme_categoryを保存するテーブル';

DROP TABLE IF EXISTS post_category;
DROP TABLE IF EXISTS tourist_spot_category;
DROP TABLE IF EXISTS vlog_category;
DROP TABLE IF EXISTS hashtag_category;

DROP TABLE IF EXISTS category;

ALTER TABLE lcategory ADD slug VARCHAR(255) NOT NULL COMMENT 'スラグ' AFTER name;

