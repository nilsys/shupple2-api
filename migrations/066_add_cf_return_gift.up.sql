CREATE TABLE IF NOT EXISTS cf_return_gift (
  id                 BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  cf_project_id      BIGINT UNSIGNED NOT NULL COMMENT '紐づくクラウドファンディングプロジェクトid',
  thumbnail          VARCHAR(255) NOT NULL,
  sort_order         INT UNSIGNED NOT NULL COMMENT '1以上で昇順',
  gift_type          TINYINT         NOT NULL COMMENT 'アクション対象',
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at         DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT cf_return_gift_cf_project_id FOREIGN KEY(cf_project_id) REFERENCES cf_project(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;