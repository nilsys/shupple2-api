CREATE TABLE IF NOT EXISTS shipping_address (
  id                BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id           BIGINT UNSIGNED NOT NULL,
  first_name        VARCHAR(255) NOT NULL COMMENT '姓',
  last_name         VARCHAR(255) NOT NULL COMMENT '名',
  first_name_kana   VARCHAR(255) NOT NULL COMMENT '姓カナ',
  last_name_kana   VARCHAR(255) NOT NULL COMMENT '名カナ',
  phone_number      VARCHAR(255) NOT NULL COMMENT '電話番号',
  postal_number     VARCHAR(255) NOT NULL COMMENT '郵便番号',
  prefecture        VARCHAR(255) NOT NULL COMMENT '都道府県',
  city              VARCHAR(255) NOT NULL COMMENT '市区町村',
  address           VARCHAR(255) NOT NULL COMMENT '番地　',
  building          VARCHAR(255) COMMENT '建物',
  email             VARCHAR(255) NOT NULL,
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at        DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT shipping_address_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = '';