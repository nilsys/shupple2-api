CREATE TABLE IF NOT EXISTS report (
  id            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id       BIGINT UNSIGNED NOT NULL COMMENT '通報したユーザーID',
  target_id     BIGINT UNSIGNED NOT NULL COMMENT '通報対象ID',
  target_type   TINYINT         NOT NULL COMMENT '通報対象タイプ',
  reason        TINYINT         NOT NULL COMMENT '通報理由タイプ',
  is_done       TINYINT         NOT NULL COMMENT '対応フラグ',
  created_at    DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  CONSTRAINT report_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = '通報を保存するテーブル'