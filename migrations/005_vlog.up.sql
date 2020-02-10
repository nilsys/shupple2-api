CREATE TABLE IF NOT EXISTS vlog (
  id          BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  user_id     BIGINT UNSIGNED NOT NULL COMMENT '執筆者',
  slug        VARCHAR(255) NOT NULL COMMENT 'wordpressのslugを入れる',
  title       VARCHAR(1024) NOT NULL COMMENT '記事タイトル',
  body        LONGTEXT NOT NULL COMMENT '本文',
  youtube_url VARCHAR(255) NOT NULL COMMENT 'youtubeのURL',
  series      VARCHAR(255) NOT NULL COMMENT 'シリーズ名',
  user_sns    VARCHAR(255) NOT NULL COMMENT '投稿者のSNS',
  editor_name VARCHAR(255) NOT NULL COMMENT '編集者の名前',
  yearmonth   VARCHAR(255) NOT NULL COMMENT '撮影年月。year_monthがまさかの予約語',
  play_time   VARCHAR(255) NOT NULL COMMENT '再生時間',
  timeline    TEXT NOT NULL COMMENT '目次みたいなやつ',
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at  DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX vlog_slug(slug),
  CONSTRAINT vlog_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS vlog_category (
  vlog_id     BIGINT UNSIGNED NOT NULL,
  category_id BIGINT UNSIGNED NOT NULL,
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(vlog_id, category_id),
  CONSTRAINT vlog_category_vlog_id FOREIGN KEY(vlog_id) REFERENCES vlog(id),
  CONSTRAINT vlog_category_category_id FOREIGN KEY(category_id) REFERENCES category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'vlogに紐づくcategoryを保存するテーブル';

CREATE TABLE IF NOT EXISTS vlog_tourist_spot (
  vlog_id         BIGINT UNSIGNED NOT NULL,
  tourist_spot_id BIGINT UNSIGNED NOT NULL,
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at      DATETIME DEFAULT NULL,
  PRIMARY KEY(vlog_id, tourist_spot_id),
  CONSTRAINT vlog_tourist_spot_vlog_id FOREIGN KEY(vlog_id) REFERENCES vlog(id),
  CONSTRAINT vlog_tourist_spot_tourist_spot_id FOREIGN KEY(tourist_spot_id) REFERENCES tourist_spot(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'vlogに紐づくtourist_spotを保存するテーブル';
