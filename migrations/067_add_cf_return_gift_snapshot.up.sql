CREATE TABLE IF NOT EXISTS cf_return_gift_snapshot (
  id                BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  cf_return_gift_id BIGINT UNSIGNED NOT NULL,
  thumbnail         VARCHAR(255) NOT NULL,
  body              VARCHAR(1024) NOT NULL COMMENT '内容',
  price             BIGINT UNSIGNED NOT NULL,
  full_amount       BIGINT UNSIGNED NOT NULL COMMENT '総数',
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at        DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT cf_return_gift_snapshot_cf_return_gift_id FOREIGN KEY(cf_return_gift_id) REFERENCES cf_return_gift(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'return_giftの商品情報テーブル、更新がある度に追加されていく';