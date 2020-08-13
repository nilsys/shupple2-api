CREATE TABLE IF NOT EXISTS user_sales_history (
  user_id            BIGINT UNSIGNED NOT NULL COMMENT '売り上げを持つuser.id',
  payment_id         BIGINT UNSIGNED COMMENT '売り上げに変更を起こしたpayment.id(入金申請等複数の場合がある為null許容)',
  deposit_request_id BIGINT UNSIGNED COMMENT 'トリガーが入金申請の場合deposit_request.idが入る',
  reason             TINYINT NOT NULL COMMENT '売り上げに変更が起きたトリガー',
  price              BIGINT NOT NULL COMMENT '変更金額、マイナスも許容(例：キャンセルされた時に売り上げからマイナス1000円)',
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT user_sales_history_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT user_sales_history_payment_id FOREIGN KEY(payment_id) REFERENCES payment(id),
  CONSTRAINT user_sales_history_deposit_request_id FOREIGN KEY(deposit_request_id) REFERENCES deposit_request(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'userの売り上げ履歴';
