CREATE TABLE IF NOT EXISTS post (
  id                BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  user_id           BIGINT UNSIGNED NOT NULL COMMENT '執筆者',
  title             VARCHAR(1024) NOT NULL COMMENT '記事タイトル',
  toc               TEXT NOT NULL COMMENT '記事の目次(jsonで格納)',
  favorite_count    INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'いいね数',
  facebook_count    INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'facebookシェア数',
  twitter_count     INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'twitterシェア数',
  slug              VARCHAR(255) NOT NULL COMMENT 'wordpressのslugを入れる',
  views             BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '閲覧数',
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at        DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX post_slug(slug),
  FULLTEXT KEY(title) WITH PARSER NGRAM,
  CONSTRAINT post_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'サイムネイルURLはIDから生成する';

CREATE TABLE IF NOT EXISTS post_body (
  post_id    BIGINT UNSIGNED NOT NULL COMMENT '紐づく投稿のID',
  page       INT UNSIGNED NOT NULL COMMENT 'ページ数(1開始)',
  body       LONGTEXT NOT NULL COMMENT '記事の本文(htmlが入る)',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (post_id, page),
  FULLTEXT KEY(body) WITH PARSER NGRAM,
  CONSTRAINT post_body_post_id FOREIGN KEY(post_id) REFERENCES post(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS category (
  id                         BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのIDをそのまま用いる',
  parent_id                  BIGINT UNSIGNED DEFAULT NULL COMMENT 'wordpressのIDをそのまま用いる',
  name                       VARCHAR(255) NOT NULL COMMENT '名前',
  type                       TINYINT NOT NULL COMMENT 'カテゴリ種別',
  created_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at                 DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  FULLTEXT KEY(name) WITH PARSER NGRAM,
  CONSTRAINT category_parent_id FOREIGN KEY(parent_id) REFERENCES category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'wordpressのcategoryに相当する';

CREATE TABLE IF NOT EXISTS post_category (
  post_id     BIGINT UNSIGNED NOT NULL,
  category_id BIGINT UNSIGNED NOT NULL COMMENT 'wordpressのcategoryに相当する',
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(post_id, category_id),
  CONSTRAINT post_category_post_id FOREIGN KEY(post_id) REFERENCES post(id),
  CONSTRAINT post_category_category_id FOREIGN KEY(category_id) REFERENCES category(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'postに紐づくcategoryを保存するテーブル';

CREATE TABLE hashtag (
  id         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  name       VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY(id),
  FULLTEXT KEY(name) WITH PARSER NGRAM
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS post_hashtag (
  post_id    BIGINT UNSIGNED NOT NULL,
  hashtag_id BIGINT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY(post_id, hashtag_id),
  CONSTRAINT post_hashtag_post_id FOREIGN KEY(post_id) REFERENCES post(id),
  CONSTRAINT post_hashtag_hashtag_id FOREIGN KEY(hashtag_id) REFERENCES hashtag(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'postに紐づくhashtagを保存するテーブル';
