CREATE TABLE IF NOT EXISTS cf_project (
  id                 BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  user_id            BIGINT UNSIGNED NOT NULL COMMENT '執筆者',
  latest_snapshot_id BIGINT UNSIGNED DEFAULT NULL COMMENT '最新のcf_profject_snapshot.id。正常なデータであればNULLにはならない',
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at         DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT cf_project_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = '更新がある場合はcf_project_snapshotに1レコード挿入される';

CREATE TABLE IF NOT EXISTS cf_project_snapshot (
  id            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  cf_project_id BIGINT UNSIGNED NOT NULL,
  user_id       BIGINT UNSIGNED NOT NULL COMMENT '執筆者',
  title         VARCHAR(1024) NOT NULL COMMENT 'プロジェクト名',
  summary       TEXT NOT NULL COMMENT 'サマリ',
  thumbnail     VARCHAR(255) NOT NULL,
  body          LONGTEXT NOT NULL COMMENT '記事の本文(htmlが入る)',
  goal_price    INT UNSIGNED NOT NULL COMMENT '目標金額',
  deadline      DATETIME NOT NULL COMMENT '締切日',
  is_attention  TINYINT NOT NULL DEFAULT 0 COMMENT '注目プロジェクトとして表示するかどうか',
  created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at    DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT cf_project_snapshot_cf_profject_id FOREIGN KEY(cf_project_id) REFERENCES cf_project(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'cf_projectの内容';

ALTER TABLE cf_project ADD CONSTRAINT cf_project_latest_snapshot_id FOREIGN KEY(latest_snapshot_id) REFERENCES cf_project_snapshot(id);

CREATE TABLE IF NOT EXISTS cf_project_snapshot_area_category (
  cf_project_snapshot_id BIGINT UNSIGNED NOT NULL,
  area_category_id BIGINT UNSIGNED NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (cf_project_snapshot_id,area_category_id),
  KEY cf_project_snapshot_area_category_area_category_id (area_category_id),
  CONSTRAINT cf_project_snapshot_area_category_area_category_id FOREIGN KEY (area_category_id) REFERENCES area_category (id),
  CONSTRAINT cf_project_snapshot_area_category_cf_project_snapshot_id FOREIGN KEY (cf_project_snapshot_id) REFERENCES cf_project_snapshot (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='cf_project_snapshotに紐づくarea_categoryを保存するテーブル';

CREATE TABLE IF NOT EXISTS cf_project_snapshot_theme_category (
  cf_project_snapshot_id BIGINT UNSIGNED NOT NULL,
  theme_category_id BIGINT UNSIGNED NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (cf_project_snapshot_id,theme_category_id),
  KEY cf_project_snapshot_theme_category_theme_category_id (theme_category_id),
  CONSTRAINT cf_project_snapshot_theme_category_theme_category_id FOREIGN KEY (theme_category_id) REFERENCES theme_category (id),
  CONSTRAINT cf_project_snapshot_theme_category_cf_project_snapshot_id FOREIGN KEY (cf_project_snapshot_id) REFERENCES cf_project_snapshot (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='cf_project_snapshotに紐づくtheme_categoryを保存するテーブル';