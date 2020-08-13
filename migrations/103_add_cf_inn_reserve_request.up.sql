CREATE TABLE IF NOT EXISTS cf_inn_reserve_request (
  user_id            BIGINT UNSIGNED NOT NULL COMMENT '予約user.id',
  payment_id         BIGINT UNSIGNED NOT NULL COMMENT 'payment.id',
  cf_return_gift_id  BIGINT UNSIGNED NOT NULL COMMENT 'payment_cf_return_gift.cf_return_gift_id',
  first_name         VARCHAR(255) NOT NULL COMMENT '名',
  last_name          VARCHAR(255) NOT NULL COMMENT '姓',
  first_name_kana    VARCHAR(255) NOT NULL COMMENT '名カナ',
  last_name_kana     VARCHAR(255) NOT NULL COMMENT '姓カナ',
  email              VARCHAR(255) NOT NULL COMMENT 'email',
  phone_number       VARCHAR(255) NOT NULL COMMENT '電話番号',
  checkin_at         DATETIME NOT NULL COMMENT 'チェックイン日時',
  checkout_at        DATETIME NOT NULL COMMENT 'チェックアウト日時',
  stay_days          BIGINT UNSIGNED NOT NULL COMMENT '宿泊日数',
  adult_member_count BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '大人人数',
  child_member_count BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '子供人数',
  remark             VARCHAR(1024) COMMENT '備考',
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT cf_inn_reserve_request_user_id FOREIGN KEY(user_id) REFERENCES user(id),
  CONSTRAINT cf_inn_reserve_request_payment_id FOREIGN KEY(payment_id) REFERENCES payment(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT = 'userの売り上げ履歴';
