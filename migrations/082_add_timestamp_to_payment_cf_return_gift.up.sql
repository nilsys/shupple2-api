ALTER TABLE cf_return_gift_snapshot
    ADD delivery_date DATETIME NOT NULL COMMENT 'お届け日時' AFTER deadline;

ALTER TABLE payment_cf_return_gift
    ADD owner_confirmed_at DATETIME DEFAULT NULL COMMENT 'オーナー対応日時' AFTER is_owner_confirmed;