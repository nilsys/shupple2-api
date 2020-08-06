ALTER TABLE payment
    ADD commission_price BIGINT UNSIGNED NOT NULL COMMENT '手数料' AFTER total_price;