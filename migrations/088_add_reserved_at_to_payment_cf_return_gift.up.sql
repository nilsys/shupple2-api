ALTER TABLE payment_cf_return_gift
    ADD user_reserve_requested_at DATETIME DEFAULT NULL COMMENT 'ユーザー(購入者)予約リクエスト発行日時、宿泊券の場合' AFTER owner_confirmed_at;
