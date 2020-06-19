CREATE TABLE IF NOT EXISTS payment_cf_return_gift (
  payment_id                 BIGINT UNSIGNED NOT NULL,
  cf_project_id              BIGINT UNSIGNED NOT NULL,
  cf_project_snapshot_id     BIGINT UNSIGNED NOT NULL,
  cf_return_gift_id          BIGINT UNSIGNED NOT NULL,
  cf_return_gift_snapshot_id BIGINT UNSIGNED NOT NULL,
  amount                     BIGINT UNSIGNED NOT NULL COMMENT '個数',
  created_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at                 DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at                 DATETIME DEFAULT NULL,
  PRIMARY KEY (payment_id, cf_return_gift_id),
  CONSTRAINT payment_cf_return_gift_payment_id FOREIGN KEY(payment_id) REFERENCES payment(id),
  CONSTRAINT payment_cf_project_cf_project_id FOREIGN KEY(cf_project_id) REFERENCES cf_project(id),
  CONSTRAINT payment_cf_project_cf_project_snapshot_id FOREIGN KEY(cf_project_snapshot_id) REFERENCES cf_project_snapshot(id),
  CONSTRAINT payment_cf_return_gift_cf_return_gift_id FOREIGN KEY(cf_return_gift_id) REFERENCES cf_return_gift(id),
  CONSTRAINT payment_cf_return_gift_cf_return_gift_snapshot_id FOREIGN KEY(cf_return_gift_snapshot_id) REFERENCES cf_return_gift_snapshot(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = '購入情報と商品の中間テーブル';