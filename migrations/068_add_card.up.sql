CREATE TABLE IF NOT EXISTS card (
  id                     BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id                BIGINT UNSIGNED NOT NULL,
  card_id                VARCHAR(255) NOT NULL COMMENT 'Pay.jp側のカードID',
  created_at             DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at             DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at             DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT credit_card_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'クレジットカード';