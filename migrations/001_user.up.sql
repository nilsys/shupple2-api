CREATE TABLE IF NOT EXISTS user (
  id                      BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  firebase_id             VARCHAR(255) NOT NULL COMMENT 'firebaseから返るsub',
  name                    VARCHAR(255) NOT NULL,
  email                   VARCHAR(255) NOT NULL,
  birthdate               DATE NOT NULL,
  gender                  TINYINT NOT NULL,
  prefecture              TINYINT NOT NULL,
  matching_reason         TINYINT NOT NULL,
  profile                 TEXT NOT NULL,
  latest_matching_user_id BIGINT UNSIGNED DEFAULT NULL COMMENT 'マッチング中のユーザーid',
  created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at   DATETIME DEFAULT NULL,
  PRIMARY KEY(id),
  CONSTRAINT user_latest_matching_user_id FOREIGN KEY(latest_matching_user_id) REFERENCES user(id),
  INDEX idx_user_firebase_id(firebase_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS user_image (
  uuid             CHAR(36) NOT NULL COMMENT 's3上の画像のUUID',
  user_id          BIGINT UNSIGNED NOT NULL,
  priority         INT UNSIGNED NOT NULL,
  mime_type        VARCHAR(255) NOT NULL COMMENT 'MIME type',
  created_at       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(uuid),
  CONSTRAINT user_image_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS user_matching_history (
  user_id                               BIGINT UNSIGNED NOT NULL,
  matching_user_id                      BIGINT UNSIGNED NOT NULL,
  matching_reason                       TINYINT NOT NULL,
  user_main_matching_approve            TINYINT DEFAULT NULL,
  matching_user_main_matching_approve   TINYINT DEFAULT NULL,
  matching_expired_at                   DATETIME NOT NULL COMMENT 'マッチング有効期限',
  main_matching_status                  TINYINT DEFAULT NULL
  created_at                            DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at                            DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(user_id, matching_user_id),
  CONSTRAINT user_matching_history_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_matching_history_matching_user_id FOREIGN KEY(matching_user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
