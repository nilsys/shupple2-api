ALTER TABLE payment
    ADD project_owner_id BIGINT UNSIGNED NOT NULL COMMENT 'オーナー(user)id' AFTER user_id,
    ADD CONSTRAINT payment_projec_owner_id FOREIGN KEY(project_owner_id) REFERENCES user(id),
    ADD total_price BIGINT UNSIGNED NOT NULL COMMENT '総金額' AFTER shipping_address_id;

ALTER TABLE payment_cf_return_gift
    ADD is_canceled TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'キャンセルされたか' AFTER amount,
    ADD is_owner_confirmed TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'オーナー側対応済か' AFTER is_canceled;
