CREATE TABLE IF NOT EXISTS deposit_request (
  id                 BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id            BIGINT UNSIGNED NOT NULL COMMENT '申請したuser(owner).id',
  price              BIGINT UNSIGNED NOT NULL COMMENT '申請金額',
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  CONSTRAINT deposit_request_user_id FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'user(owner)の売り上げ入金申請';

CREATE TABLE IF NOT EXISTS deposit_request_payment (
  deposit_request_id BIGINT UNSIGNED NOT NULL COMMENT '入金申請deposit_request.id',
  payment_id         BIGINT UNSIGNED NOT NULL COMMENT '入金申請に含まれるpayment.id',
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(deposit_request_id, payment_id),
  CONSTRAINT deposit_request_payment_deposit_request_id FOREIGN KEY(deposit_request_id) REFERENCES deposit_request(id),
  CONSTRAINT deposit_request_payment_payment_id FOREIGN KEY(payment_id) REFERENCES payment(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'deposit_requestとpaymentの中間テーブル';