CREATE TABLE IF NOT EXISTS payment (
  id                     BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id                BIGINT UNSIGNED NOT NULL,
  card_id                BIGINT UNSIGNED NOT NULL COMMENT '決済したカードID',
  charge_id              VARCHAR(255) NOT NULL COMMENT 'pay.jp側の支払いID',
  shipping_address_id    BIGINT UNSIGNED NOT NULL,
  created_at             DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at             DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at             DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT payment_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT payment_card_id FOREIGN KEY(card_id) REFERENCES card(id),
  CONSTRAINT payment_shipping_address_id FOREIGN KEY(shipping_address_id) REFERENCES shipping_address(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = '購入情報';