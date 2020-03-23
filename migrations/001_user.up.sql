CREATE TABLE IF NOT EXISTS user (
  id           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  cognito_id   VARCHAR(255) NOT NULL COMMENT 'cognitoから返るsub',
  wordpress_id INT UNSIGNED DEFAULT NULL COMMENT 'wordperssでのid',
  name         VARCHAR(255) NOT NULL,
  email        VARCHAR(255) NOT NULL,
  birthdate    DATE NOT NULL,
  gender       TINYINT NOT NULL,
  profile      TEXT NOT NULL,
  avatar_uuid  CHAR(36) NOT NULL COMMENT 's3上のアバター画像のUUID',
  created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at   DATETIME DEFAULT NULL,
  PRIMARY KEY(id),
  INDEX idx_user_cognito_id(cognito_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
