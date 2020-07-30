CREATE TABLE IF NOT EXISTS user_attribute (
  user_id                BIGINT UNSIGNED NOT NULL,
  attribute              TINYINT NOT NULL DEFAULT 0,
  created_at             DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at             DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, attribute),
  CONSTRAINT user_attribute_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'Userがどの属性を持つか';